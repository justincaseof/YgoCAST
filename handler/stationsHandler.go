package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ygost/model"
)

func StationsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("StationsHandler: ", request.RequestURI)

	result, err := xml.Marshal(model.MyStationsItems["Chillout"])
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)

}
