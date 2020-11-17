package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"ygost/handler"
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
