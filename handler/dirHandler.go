package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ygost/model"
)

func DirHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("DirHandler: ", request.RequestURI)

	result, err := xml.Marshal(model.MyStations)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)
}
