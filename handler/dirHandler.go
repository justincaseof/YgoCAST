package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ygost/helper"
)

func DirHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("DirHandler")
	items := helper.ParseFile("_examples/01_my_stations.xml")

	result, err := xml.Marshal(items)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)
}
