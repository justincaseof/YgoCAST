package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ygost/helper"
)

func StationsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("StationsHandler")
	items := helper.ParseFile("_examples/02_my_stations-GROUPNAME.xml")

	result, err := xml.Marshal(items)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)

}
