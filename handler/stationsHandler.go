package handler

import (
	"fmt"
	"net/http"
)

func StationsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("StationsHandler")
}
