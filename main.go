package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"ygost/handler"
	"ygost/helper"
	"ygost/model"
)

func main() {
	fmt.Println("Hi!")

	loadData()
	addHandlers()
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

func addHandlers() {
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/my_stations", handler.DirHandler)
	http.HandleFunc("/my_stations/", handler.StationsHandler)
}

func startServer() {
	s := &http.Server{
		Addr:           ":80",
		Handler:        nil, // if nil, uses http.DefaultServeMux
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
