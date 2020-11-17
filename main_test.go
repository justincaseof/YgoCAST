package main

import "testing"

func TestParseSuccessOfRoot(t *testing.T) {
	ParseFile("_examples/00_root.xml")
	ParseFile("_examples/01_my_stations.xml")
	ParseFile("_examples/02_my_stations-GROUPNAME.xml")
}
