package main

import (
	"flag"
	"fmt"
	"github.com/felixge/httpsnoop"
	"log"
	"net/http"
	"time"
	"ygost/handler"
	"ygost/helper"
	"ygost/middleware"
	"ygost/model"
)

// flag vars
var stationsFile string
var listenPort string

// _________

func main() {
	fmt.Println("######################################")
	fmt.Println("#               YgoCAST              #")
	fmt.Println("#              ~ 0.0.1 ~             #")
	fmt.Println("# ---------------------------------- #")
	fmt.Println("# easy Webradio Index                #")
	fmt.Println("# for YAMAHA RX-V*** receivers       #")
	fmt.Println("# github.com/justincaseof/YgoCAST    #")
	fmt.Println("# 2020-11-22                         #")
	fmt.Println("######################################")

	flag.StringVar(&stationsFile,
		"stations",
		"my_stations.yaml",
		"e.g. -stations=my_stations.yaml --> Path of a stations file (defaults to 'stations.yaml')")
	flag.StringVar(&listenPort,
		"port",
		"80",
		"e.g. -port=8087 --> the port for the server to listen on")
	flag.Parse()

	loadStations(stationsFile)
	startServer(listenPort)
}

// our source is a yaml file.
func loadStations(fileName string) {
	model.STATIONS = helper.ParseYaml(fileName)

	model.STATIONS_BY_ID = make(map[string]*model.StationInfo)
	// populate ids and parent dirs
	fmt.Printf("Populating stations...\n")
	for _, dir := range model.STATIONS.SubDirectoriesAsList() {
		fmt.Printf("  [%s]:\n", dir.Name)
		for _, sta := range dir.Stations {
			id := sta.GenerateStationID(dir.Name)
			if model.STATIONS_BY_ID[id] == nil {
				fmt.Printf("    * adding station '%s': '%s'\n", id, sta.Name)
				sta.ParentDirName = dir.Name
				sta.StationId = id
				model.STATIONS_BY_ID[id] = sta
			} else {
				panic("duplicate station name: " + sta.Name)
			}
		}
	}
}

func startServer(listenPort string) {
	// add xml header XML_HEADER = '<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>'

	mux := &http.ServeMux{}
	mux.HandleFunc("/ycast/icon", handler.IconHandler)
	mux.HandleFunc("/icon", handler.IconHandler)
	mux.HandleFunc("/setupapp/Yamaha/asp/BrowseXML/statxml.asp", handler.SetupHandlerStat)
	mux.HandleFunc("/setupapp/Yamaha/asp/BrowseXML/loginXML.asp", handler.SetupHandlerLogin)
	mux.HandleFunc("/ycast/my_stations/", handler.StationsHandler)
	mux.HandleFunc("/ycast/radiobrowser/", handler.RadiobrowserHandler)
	mux.HandleFunc("/ycast", handler.RootHandler)
	mux.HandleFunc("/", handler.RootHandler)

	// wrap mux with our logger. this will
	var loggingWrap http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(mux, w, r)

		log.Printf(
			"%s %s (code=%d dt=%s written=%d)",
			r.Method,
			r.URL,
			m.Code,
			m.Duration,
			m.Written,
		)
	})
	// ... potentially add more middleware handlers

	s := &http.Server{
		//Addr:           ":80",
		Addr:           ":" + listenPort,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        middleware.XMLEncodingLineAddingWrapper(middleware.StationIdMiddleware(middleware.MACAddressMiddleware(loggingWrap))),
	}
	log.Fatal(s.ListenAndServe())
}
