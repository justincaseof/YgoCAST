package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"ygost/model"
)

func StationsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("StationsHandler")

	if strings.Contains(request.RequestURI, "/?vtuner") {
		// we need to respond with our genres/directories
		fmt.Println("  --> responding Directories")
		result, err := xml.Marshal(model.MyStations)
		if err != nil {
			fmt.Println("cannot marshall")
		}
		writer.Write(result)
	} else {
		fmt.Println("  --> responding Stations of Directory")

		// TODO: extract Directory name

		directoryName := "Chillout" // FIXME dummy
		result, err := xml.Marshal(model.MyStationsItems[directoryName])
		if err != nil {
			fmt.Println("cannot marshall")
		}
		writer.Write(result)
	}

}

func StationById(id string) model.ListOfItems {
	// TODO: get actual station by id
	return model.ListOfItems{
		ItemCount: 1,
		Items: []model.Item{
			{
				ItemType:      model.Station,
				StationId:     id,
				StationName:   "dummy FooBar",
				StationUrl:    "http://ice1.somafm.com/dronezone-256-mp3",
				StationDesc:   "Chillout",
				StationFormat: "Chillout",
				Logo:          "http://192.168.178.61/ycast/logo?id=MY_B08E60982371",
			},
		},
	}
}
