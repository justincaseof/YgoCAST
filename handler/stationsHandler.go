package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"time"
	"ygost/model_yamaha"
)

func RadiobrowserHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RadiobrowserHandler")

	//id := middleware.GenerateStationID()
	dummy := model_yamaha.ListOfItems{
		ItemCount: 1,
		Items: []model_yamaha.Item{
			model_yamaha.Item{
				ItemType: model_yamaha.Dir,
				//Title: id,
				Title:        time.Now().Format("15:04:05") + "." + fmt.Sprintf("%d", time.Now().Nanosecond()),
				UrlDir:       "http://192.168.178.61/ycast/radiobrowser/dummy", // FIXME: IP
				UrlDirBackUp: "http://192.168.178.61/ycast/radiobrowser/dummy", // FIXME: IP
				DirCount:     1,
			},
		},
	}

	result, err := xml.Marshal(dummy)
	if err != nil {
		fmt.Println("cannot marshall")
	}
	writer.Write(result)
}

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
		result, err := xml.Marshal(model_yamaha.YamahaMyStations)
		if err != nil {
			fmt.Println("cannot marshall")
		}
		writer.Write(result)
	} else {
		fmt.Println("  --> responding Stations of Directory ", directoryName)
		subdirList := model_yamaha.MyStationsDirNameToListOfItemsMapping[directoryName]
		subdirList.ItemCount = int32(len(subdirList.Items)) // update, just in case

		result, err := xml.Marshal(subdirList)
		if err != nil {
			fmt.Println("cannot marshall")
		}
		writer.Write(result)
	}

}

func SingleStationById(id string) model_yamaha.ListOfItems {
	result := model_yamaha.ListOfItems{
		ItemCount: 1,
		Items: []model_yamaha.Item{
			model_yamaha.Item(model_yamaha.StationIDtoStationMapping[id]),
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
