package handler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"ygost/model"
	"ygost/model_yamaha"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RootHandler", request.RequestURI)

	baseUrl := "http://" + request.Host + "/ycast/" // FIXME -->config

	root := model_yamaha.RootList{
		Items: []model_yamaha.Item{},
	}
	root.Items = append(root.Items, model_yamaha.Item(model_yamaha.DirectoryItem{
		ItemType:     model_yamaha.Dir,
		Title:        "My Stations",
		UrlDir:       baseUrl + "my_stations/?vtuner=true",
		UrlDirBackUp: baseUrl + "my_stations/?vtuner=true",
		DirCount:     int32(len(model.STATIONS.SubDirectories)),
	}))
	root.Items = append(root.Items, model_yamaha.Item(model_yamaha.DirectoryItem{
		ItemType:     model_yamaha.Dir,
		Title:        "### Gruss vom Tobse :-) ###",
		UrlDir:       baseUrl + "radiobrowser/?vtuner=true",
		UrlDirBackUp: baseUrl + "radiobrowser/?vtuner=true",
		DirCount:     1,
	}))
	//root.ItemCount = int32(len(root.Items))
	root.ItemCount = -1 // needs to be -1 (i guess)

	result, err := xml.Marshal(root)
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

	baseUrl := "http://" + request.Host + "/ycast/" // FIXME -->config

	writer.Write(SingleStationById(fmt.Sprintf("%s", stationId), baseUrl).MarshalToXML())
}

func SetupHandlerLogin(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("SetupHandlerLogin")

	// 2 Types:
	//  * /setupapp/Yamaha/asp/BrowseXML/loginXML.asp?token=0
	//    --> simply respond with a token (even a dummy) <EncryptedToken>0000000000000000</EncryptedToken>
	//    --> !!! ATTENTION: RX-V receiver will REJECT the response if it contains '<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>' reamble in response!
	//  * /setupapp/Yamaha/asp/BrowseXML/loginXML.asp?mac=e3629f8b2113402738c4d17f406793dc&dlang=eng&fver=W&start=1&howmany=8
	//    --> simply respond with YamahaRoot
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

	fileNames := []string{
		"_examples/icons/icon1.jpeg",
		"_examples/icons/icon2.jpeg",
		"_examples/icons/icon3.jpeg",
		"_examples/icons/icon4.jpeg",
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Uint64() % uint64(len(fileNames))

	iconFile, err := os.Open(fileNames[idx])
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
