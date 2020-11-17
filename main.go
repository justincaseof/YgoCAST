package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"ygost/handler"
	"ygost/model"
)

func main() {
	fmt.Println("Hi!")

	addHandlers()
	startServer()
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

func ParseFile(name string) {
	// ##################################################################################
	xmlFile, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened file " + name)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// ##################################################################################
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var listOfItems model.ListOfItems
	xml.Unmarshal(byteValue, &listOfItems)

	fmt.Println("done.", listOfItems)
}
