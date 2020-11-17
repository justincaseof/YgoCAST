package model

import "encoding/xml"

type ListOfItems struct {
	XMLName   xml.Name `xml:"ListOfItems"`
	ItemCount int32    `xml:"ItemCount"`
	Items     []Item   `xml:"Item"`
}

type Item struct {
	XMLName  xml.Name `xml:"Item"`
	ItemType string   `xml:"ItemType"`

	// ItemType="Dir"
	Title        string `xml:"Title"`
	UrlDir       string `xml:"UrlDir"`
	UrlDirBackUp string `xml:"UrlDirBackUp"`
	DirCount     int32  `xml:"DirCount"`

	// ItemType="Station"
	StationId        string `xml:"StationId"`
	StationName      string `xml:"StationName"`
	StationUrl       string `xml:"StationUrl"`
	StationDesc      string `xml:"StationDesc"`
	Logo             string `xml:"Logo"`
	StationFormat    string `xml:"StationFormat"`
	StationLocation  string `xml:"StationLocation"`
	StationBandWidth string `xml:"StationBandWidth"`
	StationMime      string `xml:"StationMime"`
	Relia            string `xml:"Relia"`
	Bookmark         string `xml:"Bookmark"`
}

// ###

var Root ListOfItems = ListOfItems{
	ItemCount: -1, // looks like to be some kind of default for root folder
	Items: []Item{
		Item{
			ItemType:     "Dir",
			Title:        "Radiobrowser",
			UrlDir:       "http:// TODO",
			UrlDirBackUp: "http:// TODO",
			DirCount:     4,
		},
		Item{
			ItemType:     "Dir",
			Title:        "My Stations",
			UrlDir:       "http:// TODO",
			UrlDirBackUp: "http:// TODO",
			DirCount:     4,
		},
	},
}
var MyStations ListOfItems
var Chill ListOfItems
