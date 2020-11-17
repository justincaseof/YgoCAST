package handler

import (
	"fmt"
	"net/http"
)

func DirHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("DirHandler")
}
