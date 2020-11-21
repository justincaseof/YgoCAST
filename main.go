package main

import (
	"fmt"
	"github.com/felixge/httpsnoop"
	"log"
	"net/http"
	"time"
	"ygost/handler"
	"ygost/helper"
	"ygost/middleware"
	"ygost/model"
	"ygost/model_yamaha"
)

func main() {
	fmt.Println("Hi!")

	loadData()
	startServer()
}

func loadData() {
	//IP := "192.168.178.61"
	//Path := "/"

	// ### ROOT --> still static
	model_yamaha.YamahaRoot = helper.ParseYamahaXMLFile("_examples/dev/00_root.xml")

	// ### MYSTATIONS
	//model.YamahaMyStations = helper.ParseYamahaXMLFile("_examples/dev/01-01_my_stations.xml")

	// ### My-Stations Folders
	//Jungletrain := helper.ParseYamahaXMLFile("_examples/dev/02-00_my_stations-Jungletrain.xml")
	//Electronic := helper.ParseYamahaXMLFile("_examples/dev/02-01_my_stations-Electronic.xml")
	//Chillout := helper.ParseYamahaXMLFile("_examples/dev/02-02_my_stations-Chillout.xml")
	//IntergalacticFM := helper.ParseYamahaXMLFile("_examples/dev/02-03_my_stations-IntergalacticFM.xml")
	//model.MyStationsDirNameToListOfItemsMapping = make(map[string]model.ListOfItems)
	//model.MyStationsDirNameToListOfItemsMapping["Jungletrain"] = Jungletrain
	//model.MyStationsDirNameToListOfItemsMapping["Electronic"] = Electronic
	//model.MyStationsDirNameToListOfItemsMapping["Chillout"] = Chillout
	//model.MyStationsDirNameToListOfItemsMapping["IntergalacticFM"] = IntergalacticFM

	// ===== MAPPING STUFF =====
	// generade IDs and fill hashtable for lookup
	model_yamaha.StationIDtoStationMapping = make(map[string]model_yamaha.StationItem)
	for _, listOfItems := range model_yamaha.MyStationsDirNameToListOfItemsMapping {
		for _, item := range listOfItems.Items {
			model_yamaha.StationIDtoStationMapping[item.StationId] = model_yamaha.StationItem(item)
		}
	}

	// test
	loadStations()
}

// test
func loadStations() {
	model.STATIONS = helper.ParseYaml("_examples/my_stations.yaml")
	stationsJson := helper.ParseJSON("_examples/my_stations.json")
	fmt.Println(model.STATIONS)
	fmt.Println(stationsJson)

	//var stationsOfDir model_yamaha.StationsList
	//for key, val := range STATIONS_YAML.SubDirectories {
	//	stationsOfDir.Encode(key, val.Stations)
	//}

	fmt.Println("")
}

func startServer() {
	// add xml header XML_HEADER = '<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>'

	mux := &http.ServeMux{}
	mux.HandleFunc("/ycast/icon", handler.IconHandler)
	mux.HandleFunc("/icon", handler.IconHandler)
	mux.HandleFunc("/setupapp/Yamaha/asp/BrowseXML/statxml.asp", handler.SetupHandlerStat)
	mux.HandleFunc("/setupapp/Yamaha/asp/BrowseXML/loginXML.asp", handler.SetupHandlerLogin)
	mux.Handle("/ycast/my_stations/", middleware.XMLEncodingLineAddingWrapper(http.HandlerFunc(handler.StationsHandler)))
	mux.Handle("/ycast/radiobrowser/", middleware.XMLEncodingLineAddingWrapper(http.HandlerFunc(handler.RadiobrowserHandler)))
	mux.Handle("/ycast", middleware.XMLEncodingLineAddingWrapper(http.HandlerFunc(handler.RootHandler)))
	mux.Handle("/", middleware.XMLEncodingLineAddingWrapper(http.HandlerFunc(handler.RootHandler)))

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
		Addr:           ":80",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        middleware.StationIdMiddleware(middleware.MACAddressMiddleware(loggingWrap)),
	}
	log.Fatal(s.ListenAndServe())
}
