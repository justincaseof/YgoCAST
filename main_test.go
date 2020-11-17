package main

import (
	"fmt"
	"testing"
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

}
