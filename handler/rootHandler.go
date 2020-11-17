package handler

import (
	"fmt"
	"net/http"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RootHandler")
}
