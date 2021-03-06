package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
	"ygost/model"
	"ygost/model_yamaha"
)

func RadiobrowserHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RadiobrowserHandler")

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
	regex := regexp.MustCompile(`(.)*my_stations/([^\? /]*)`)
	if regex.MatchString(request.RequestURI) {
		subMatches := regex.FindStringSubmatch(request.RequestURI)
		directoryName = subMatches[2]
		fmt.Println("  --> Desired Directory: ", directoryName)
	} else {
		//fmt.Println("NO MATCH")
	}

	if directoryName == "" {
		// we need to respond with our genres/directories
		fmt.Println("  --> responding Directories")

		uriWithoutArgs := strings.Split(request.RequestURI, "?")[0]
		baseUrl := "http://" + request.Host + uriWithoutArgs
		result := model_yamaha.StationDirectoryList{}
		result = result.Encode(model.STATIONS.SubDirectoriesAsList(), baseUrl)

		writer.Write(result.MarshalToXML())
	} else {
		fmt.Println("  --> responding Stations of Directory ", directoryName)

		baseUrl := "http://" + request.Host + "/ycast/" // FIXME -->config
		result := model_yamaha.StationsList{}
		result = result.Encode(directoryName, model.STATIONS.SubDirectories[directoryName].Stations, baseUrl)

		writer.Write(result.MarshalToXML())
	}

}

func SingleStationById(id string, baseUrl string) model_yamaha.StationsList {
	station := model_yamaha.StationItem{}
	stationInfo := model.STATIONS_BY_ID[id]

	if stationInfo != nil {
		station = station.Encode(*stationInfo, baseUrl)
		fmt.Printf("  --> responding station '%s': '%s'\n", id, station.StationName)
		result := model_yamaha.StationsList{
			ItemCount: 1,
			Items: []model_yamaha.Item{
				model_yamaha.Item(station),
			},
		}
		return result
	}

	fmt.Printf("  --> station '%s' not found\n", id)
	return model_yamaha.StationsList{
		ItemCount: 1,
		Items: []model_yamaha.Item{
			model_yamaha.Item{ItemType: model_yamaha.Display, Display: "Station not found :-("},
		},
	}
}
