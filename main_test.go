package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"ygost/handler"
	"ygost/helper"
	"ygost/model"
	"ygost/model_yamaha"
)

func TestParseSuccessOfRoot(t *testing.T) {
	fmt.Println(helper.ParseYamahaXMLFile("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/00_root.xml"))
}

func TestParseSuccessOfMyStations(t *testing.T) {
	fmt.Println(helper.ParseYamahaXMLFile("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/01_my_stations.xml"))
}

func TestParseSuccessOfSubdirXYZ(t *testing.T) {
	fmt.Println(helper.ParseYamahaXMLFile("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/02_my_stations-GROUPNAME.xml"))
}

func TestXMLSerializationAndParsing(t *testing.T) {
	chill := model_yamaha.StationsList{}
	chill.Items = []model_yamaha.Item{ //make([]model_yamaha.Item, 0)
		model_yamaha.Item{
			ItemType:     "Dir",
			Title:        "Radiobrowser",
			UrlDir:       "http://192.168.178.61/ycast/radiobrowser/?vtuner=true",
			UrlDirBackUp: "http://192.168.178.61/ycast/radiobrowser/?vtuner=true",
			DirCount:     4,
		},
		model_yamaha.Item{
			ItemType:     "Dir",
			Title:        "My Stations",
			UrlDir:       "http://192.168.178.61/ycast/my_stations/?vtuner=true",
			UrlDirBackUp: "http://192.168.178.61/ycast/my_stations/?vtuner=true",
			DirCount:     4,
		},
	}

	res, _ := helper.XMLHelper{}.Marshal(chill)
	fmt.Println(string(res))
}

func TestYamlSerializationAndParsing(t *testing.T) {
	myStations := model.MyStationDirectories{
		map[string]model.Subdirectory{
			"Chill": model.Subdirectory{
				Name: "Chill",
				Stations: []*model.StationInfo{
					{StationName: "111", StationURL: "222", IconURL: "333"},
					{StationName: "444", StationURL: "555", IconURL: "666"},
				},
			},
			"Jungletrain*s!": model.Subdirectory{
				Name: "Jungletrain*s:",
				Stations: []*model.StationInfo{
					{StationName: "123", StationURL: "456", IconURL: "789"},
					{StationName: "987", StationURL: "654", IconURL: "321"},
				},
			},
		},
	}
	res, _ := helper.YamlHelper{}.Marshal(myStations)
	fmt.Println(string(res))
}

func TestRootHandler(t *testing.T) {
	rec := httptest.NewRecorder()

	// setup mock data
	root := model_yamaha.RootList(model_yamaha.ListOfItems{
		ItemCount: -1, // looks like to be some kind of default for root folder
		Items: []model_yamaha.Item{
			model_yamaha.Item{
				ItemType:     "Dir",
				Title:        "Radiobrowser",
				UrlDir:       "http://192.168.178.61/ycast/radiobrowser/?vtuner=true",
				UrlDirBackUp: "http://192.168.178.61/ycast/radiobrowser/?vtuner=true",
				DirCount:     4,
			},
			model_yamaha.Item{
				ItemType:     "Dir",
				Title:        "My Stations",
				UrlDir:       "http://192.168.178.61/ycast/my_stations/?vtuner=true",
				UrlDirBackUp: "http://192.168.178.61/ycast/my_stations/?vtuner=true",
				DirCount:     4,
			},
		},
	})
	fmt.Println("", root)

	// INVOKE !
	handler.RootHandler(rec, &http.Request{})

	// read expected data
	xmlFile, err := os.Open("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/00_root.xml")
	if err != nil {
		t.Fatalf("Cannot read test data")
	}
	defer xmlFile.Close()
	expectedContect, _ := ioutil.ReadAll(xmlFile)
	var listOfItemsExpected model_yamaha.ListOfItems
	xml.Unmarshal(expectedContect, &listOfItemsExpected)

	// read response and marshll
	responseBytes, _ := ioutil.ReadAll(rec.Body)
	var listOfItemsInResponse model_yamaha.ListOfItems
	xml.Unmarshal(responseBytes, &listOfItemsInResponse)

	if listOfItemsExpected.ItemCount != listOfItemsInResponse.ItemCount {
		t.Fail()
	}

	if len(listOfItemsExpected.Items) != 2 {
		t.Fail()
	}

	if len(listOfItemsExpected.Items) != len(listOfItemsInResponse.Items) {
		t.Fail()
	}

}
