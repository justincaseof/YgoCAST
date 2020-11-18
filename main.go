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
)

func main() {
	fmt.Println("Hi!")

	loadData()
	startServer()
}
func loadData() {
	//IP := "192.168.178.61"
	//Path := "/"

	model.Root = helper.ParseFile("_examples/dev/00_root.xml")
	model.MyStations = helper.ParseFile("_examples/dev/01_my_stations.xml")

	Jungletrain := helper.ParseFile("_examples/dev/02-00_my_stations-Jungletrain.xml")
	Electronic := helper.ParseFile("_examples/dev/02-01_my_stations-Electronic.xml")
	Chillout := helper.ParseFile("_examples/dev/02-02_my_stations-Chillout.xml")
	IntergalacticFM := helper.ParseFile("_examples/dev/02-03_my_stations-IntergalacticFM.xml")
	model.MyStationsItems = make(map[string]model.ListOfItems)
	model.MyStationsItems["Jungletrain"] = Jungletrain
	model.MyStationsItems["Electronic"] = Electronic
	model.MyStationsItems["Chillout"] = Chillout
	model.MyStationsItems["IntergalacticFM"] = IntergalacticFM

}

func startServer() {
	// add xml header XML_HEADER = '<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>'

	mux := &http.ServeMux{}
	mux.HandleFunc("/setupapp/Yamaha/asp/BrowseXML/statxml.asp", handler.SetupHandlerStat)
	mux.HandleFunc("/setupapp/Yamaha/asp/BrowseXML/loginXML.asp", handler.SetupHandlerLogin)
	mux.HandleFunc("/ycast/my_stations/", handler.StationsHandler)
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
		Addr:           ":80",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        middleware.MACAddressMiddleware(middleware.XMLEncodingLineAddingWrapper(loggingWrap)),
	}
	log.Fatal(s.ListenAndServe())
}
