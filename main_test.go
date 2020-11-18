package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
	"ygost/handler"
	"ygost/helper"
	"ygost/model"
)

func TestParseSuccessOfRoot(t *testing.T) {
	fmt.Println(helper.ParseFile("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/00_root.xml"))
}

func TestParseSuccessOfMyStations(t *testing.T) {
	fmt.Println(helper.ParseFile("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/01_my_stations.xml"))
}

func TestParseSuccessOfSubdirXYZ(t *testing.T) {
	fmt.Println(helper.ParseFile("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/02_my_stations-GROUPNAME.xml"))
}

func TestRootHandler(t *testing.T) {
	rec := httptest.NewRecorder()

	// setup mock data
	model.Root = model.ListOfItems{
		ItemCount: -1, // looks like to be some kind of default for root folder
		Items: []model.Item{
			model.Item{
				ItemType:     "Dir",
				Title:        "Radiobrowser",
				UrlDir:       "http://192.168.178.61/ycast/radiobrowser/?vtuner=true",
				UrlDirBackUp: "http://192.168.178.61/ycast/radiobrowser/?vtuner=true",
				DirCount:     4,
			},
			model.Item{
				ItemType:     "Dir",
				Title:        "My Stations",
				UrlDir:       "http://192.168.178.61/ycast/my_stations/?vtuner=true",
				UrlDirBackUp: "http://192.168.178.61/ycast/my_stations/?vtuner=true",
				DirCount:     4,
			},
		},
	}

	// INVOKE !
	handler.RootHandler(rec, nil)

	// read expected data
	xmlFile, err := os.Open("_examples/DO_NOT_CHANGE_TESTS_DEPEND_ON_IT/00_root.xml")
	if err != nil {
		t.Fatalf("Cannot read test data")
	}
	defer xmlFile.Close()
	expectedContect, _ := ioutil.ReadAll(xmlFile)
	var listOfItemsExpected model.ListOfItems
	xml.Unmarshal(expectedContect, &listOfItemsExpected)

	// read response and marshll
	responseBytes, _ := ioutil.ReadAll(rec.Body)
	var listOfItemsInResponse model.ListOfItems
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
