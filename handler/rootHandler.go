package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ygost/model"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RootHandler", request.RequestURI)

	result, err := xml.Marshal(model.Root)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)
}
