package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"ygost/handler"
	"ygost/helper"
)

func TestParseSuccessOfRoot(t *testing.T) {
	fmt.Println(helper.ParseFile("_examples/00_root.xml"))
}

func TestParseSuccessOfMyStations(t *testing.T) {
	fmt.Println(helper.ParseFile("_examples/01_my_stations.xml"))
}

func TestParseSuccessOfSubdirXYZ(t *testing.T) {
	fmt.Println(helper.ParseFile("_examples/02_my_stations-GROUPNAME.xml"))
}

func TestRootHandler(t *testing.T) {
	rec := httptest.NewRecorder()

	handler.RootHandler(rec, nil)

	xmlFile, err := os.Open("_examples/00_root.xml")
	if err != nil {
		t.Fatalf("Cannot read test data")
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	responseBytes, _ := ioutil.ReadAll(rec.Body)

	if !Equal(byteValue, responseBytes) {
		log.Fatalf("Not equal!")
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
