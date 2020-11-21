package model_yamaha

import (
	"encoding/xml"
	"ygost/model"
)

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

type RootList ListOfItems             // the one and only
type StationDirectoryList ListOfItems // directories/genres
type StationsList ListOfItems         // single stations of a directory/genre
type DirectoryItem Item               // an Item with ItemType=Dir
type StationItem Item                 // an Item with ItemType=Station

// ### Globally usabe vars
var YamahaRoot RootList

var MyStationsDirNameToListOfItemsMapping map[string]StationsList
var StationIDtoStationMapping map[string]StationItem

func (myStationDirectories StationDirectoryList) MarshalToXML() []byte {
	bytes, _ := xml.Marshal(myStationDirectories)
	return bytes
}
func (myStations StationsList) MarshalToXML() []byte {
	bytes, _ := xml.Marshal(myStations)
	return bytes
}

func (myStationDirectories StationDirectoryList) Encode(directories []model.Subdirectory) StationDirectoryList {
	loi := StationDirectoryList{
		Items: make([]Item, 0),
	}
	// SubDirectories
	for _, v := range directories {
		loi.Items = append(loi.Items, Item((&DirectoryItem{}).Encode(v)))
	}
	myStationDirectories.Items = loi.Items
	myStationDirectories.ItemCount = int32(len(loi.Items))
	return myStationDirectories
}
func (stationList StationsList) Encode(dirname string, stations []model.StationInfo) StationsList {
	los := StationsList{
		Items: make([]Item, 0),
	}
	// Stations
	for _, station := range stations {
		los.Items = append(los.Items, Item((&StationItem{}).Encode(dirname, station)))
	}
	stationList.Items = los.Items
	stationList.ItemCount = int32(len(los.Items))
	return stationList
}

func (subDirItem DirectoryItem) Encode(subDir model.Subdirectory) DirectoryItem {
	subDirItem.ItemType = Dir
	subDirItem.Title = subDir.Name
	subDirItem.UrlDir = "TODO"
	return subDirItem
}
func (stationItem StationItem) Encode(subDirName string, station model.StationInfo) StationItem {
	stationItem.ItemType = Station
	stationItem.Title = station.StationName
	stationItem.StationFormat = subDirName
	stationItem.StationDesc = subDirName
	stationItem.UrlDir = station.StationURL
	//stationItem.Logo = station.IconURL
	stationItem.Logo = "http://192.168.178.61/ycast/icon?foo=bar"
	return stationItem
}
