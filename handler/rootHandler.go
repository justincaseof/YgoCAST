package handler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func SetupHandlerStat(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("SetupHandlerStat")
	stationId := request.Context().Value("StationId")
	fmt.Println("  --> StationId: ", stationId)

	//  * /setupapp/Yamaha/asp/BrowseXML/statxml.asp?mac=e3629f8b2113402738c4d17f406793dc&fver=W&id=MY_B08E60982371
	//    --> simply respond with the station of given ID

	// for now, always return dummy
	singleStationById := SingleStationById(fmt.Sprintf("%s", stationId)) // what?

	result, err := xml.Marshal(singleStationById)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)
}

func SetupHandlerLogin(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("SetupHandlerLogin")

	// 2 Types:
	//  * /setupapp/Yamaha/asp/BrowseXML/loginXML.asp?token=0
	//    --> simply respond with a token (even a dummy) <EncryptedToken>0000000000000000</EncryptedToken>
	//  * /setupapp/Yamaha/asp/BrowseXML/loginXML.asp?mac=e3629f8b2113402738c4d17f406793dc&dlang=eng&fver=W&start=1&howmany=8
	//    --> simply respond with Root
	if strings.Contains(request.RequestURI, "token=0") {
		fmt.Println("  --> Sending EncryptedToken")
		writer.Write([]byte("<EncryptedToken>0000000000000000</EncryptedToken>"))
	} else {
		fmt.Println("  --> Redirecting to root.")
		RootHandler(writer, request)
	}
}

func IconHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("IconHandler")

	iconFile, err := os.Open("_examples/icon.jpeg")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer iconFile.Close()

	// ##################################################################################
	byteValue, _ := ioutil.ReadAll(iconFile)
	writer.Write(byteValue)
}
