package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"ygost/model"
)

func StationsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("StationsHandler")
	fmt.Println("  --> MAC: ", request.Context().Value("MAC"))

	directoryName := ""
	//regex := regexp.MustCompile(`(.)*my_stations/(.*)([/]?\?vtuner)`)
	regex := regexp.MustCompile(`(.)*my_stations/([^\? /]*)`)
	if regex.MatchString(request.RequestURI) {
		//fmt.Println("MATCH ! ")
		subMatches := regex.FindStringSubmatch(request.RequestURI)
		directoryName = subMatches[2]
		fmt.Println("  --> Desired Directory: ", directoryName)
	} else {
		//fmt.Println("NO MATCH")
	}

	if directoryName == "" {
		// we need to respond with our genres/directories
		fmt.Println("  --> responding Directories")
		result, err := xml.Marshal(model.MyStations)
		if err != nil {
			fmt.Println("cannot marshall")
		}
		writer.Write(result)
	} else {
		fmt.Println("  --> responding Stations of Directory ", directoryName)
		result, err := xml.Marshal(model.MyStationsDirNameToListOfItemsMapping[directoryName])
		if err != nil {
			fmt.Println("cannot marshall")
		}
		writer.Write(result)
	}

}

func SingleStationById(id string) model.ListOfItems {
	result := model.ListOfItems{
		ItemCount: 1,
		Items: []model.Item{
			model.StationIDtoStationMapping[id],
		},
	}
	return result
	//return model.ListOfItems{
	//	ItemCount: 1,
	//	Items: []model.Item{
	//		{
	//			ItemType:      model.Station,
	//			StationId:     id,
	//			StationName:   "dummy FooBar",
	//			StationUrl:    "http://ice1.somafm.com/dronezone-256-mp3",
	//			StationDesc:   "Chillout",
	//			StationFormat: "Chillout",
	//			Logo:          "http://192.168.178.61/ycast/logo?id=MY_B08E60982371",
	//		},
	//	},
}
