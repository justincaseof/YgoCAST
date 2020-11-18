package model

import "encoding/xml"

type ListOfItems struct {
	XMLName   xml.Name `xml:"ListOfItems"`
	ItemCount int32    `xml:"ItemCount"`
	Items     []Item   `xml:"Item"`
}

type Item struct {
	XMLName xml.Name `xml:"Item"`

	ItemType ItemType `xml:"ItemType"` // 'Dir' or 'Station'  (probably case sensitive)

	// ItemType="Dir"
	Title        string `xml:"Title,omitempty"`
	UrlDir       string `xml:"UrlDir,omitempty"`
	UrlDirBackUp string `xml:"UrlDirBackUp,omitempty"`
	DirCount     int32  `xml:"DirCount,omitempty"`

	// ItemType="Station"
	StationId        string `xml:"StationId,omitempty"`
	StationName      string `xml:"StationName,omitempty"`
	StationUrl       string `xml:"StationUrl,omitempty"`
	StationDesc      string `xml:"StationDesc,omitempty"`
	Logo             string `xml:"Logo,omitempty"`
	StationFormat    string `xml:"StationFormat,omitempty"`
	StationLocation  string `xml:"StationLocation,omitempty"`
	StationBandWidth string `xml:"StationBandWidth,omitempty"`
	StationMime      string `xml:"StationMime,omitempty"`
	Relia            string `xml:"Relia,omitempty"`
	Bookmark         string `xml:"Bookmark,omitempty"`

	// ItemType="Display"
	Display string `xml:"Display,omitempty"`
}

type ItemType string

const (
	Dir     ItemType = "Dir"
	Station ItemType = "Station"
	Display ItemType = "Display"
)

// ####################################

var Root ListOfItems = ListOfItems{
	ItemCount: -1, // looks like to be some kind of default for root folder
	Items: []Item{
		Item{
			ItemType:     "Dir",
			Title:        "Radiobrowser", // dont't change. works as 'id' and will be key for our tree
			UrlDir:       "http:// TODO",
			UrlDirBackUp: "http:// TODO",
			DirCount:     4,
		},
		Item{
			ItemType:     "Dir",
			Title:        "My Stations", // dont't change. works as 'id' and will be key for our tree
			UrlDir:       "http:// TODO",
			UrlDirBackUp: "http:// TODO",
			DirCount:     4,
		},
	},
}

var MyStations ListOfItems

var MyStationsItems map[string]ListOfItems
