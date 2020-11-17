package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ygost/helper"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RootHandler")
	items := helper.ParseFile("_examples/00_root.xml")

	result, err := xml.Marshal(items)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)
}
