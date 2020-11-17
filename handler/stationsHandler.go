package handler

import (
	"fmt"
	"net/http"
	"ygost/helper"
)

func StationsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("StationsHandler")
	helper.ParseFile("02_my_stations-GROUPNAME.xml")

}
