package model_yamaha

import (
	"encoding/xml"
	"strings"
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

func (myStationDirectories StationDirectoryList) MarshalToXML() []byte {
	bytes, _ := xml.Marshal(myStationDirectories)
	return bytes
}
func (myStations StationsList) MarshalToXML() []byte {
	bytes, _ := xml.Marshal(myStations)
	return bytes
}

func (myStationDirectories StationDirectoryList) Encode(directories []model.Subdirectory, requestUrl string) StationDirectoryList {
	loi := StationDirectoryList{
		Items: make([]Item, 0),
	}
	// SubDirectories
	for _, v := range directories {
		loi.Items = append(loi.Items, Item((&DirectoryItem{}).Encode(v, requestUrl)))
	}
	myStationDirectories.Items = loi.Items
	myStationDirectories.ItemCount = int32(len(loi.Items))
	return myStationDirectories
}
func (stationList StationsList) Encode(dirname string, stations []*model.StationInfo, baseUrl string) StationsList {
	los := StationsList{
		Items: make([]Item, 0),
	}
	// Stations
	for _, station := range stations {
		station.ParentDirName = dirname
		los.Items = append(los.Items, Item((&StationItem{}).Encode(*station, baseUrl)))
	}
	stationList.Items = los.Items
	stationList.ItemCount = int32(len(los.Items))
	return stationList
}

func (subDirItem DirectoryItem) Encode(subDir model.Subdirectory, baseUrl string) DirectoryItem {
	subDirItem.ItemType = Dir
	subDirItem.Title = subDir.Name
	subDirItem.UrlDir = baseUrl + subDir.Name + "?vtuner=true" // FIXME yikes!
	subDirItem.UrlDirBackUp = subDirItem.UrlDir
	subDirItem.DirCount = int32(len(subDir.Stations))
	return subDirItem
}
func (stationItem StationItem) Encode(station model.StationInfo, baseUrl string) StationItem {
	stationItem.ItemType = Station
	stationItem.StationId = station.StationId
	stationItem.StationName = station.StationName
	stationItem.StationFormat = station.ParentDirName
	stationItem.StationDesc = station.ParentDirName
	stationItem.StationUrl = station.StationURL
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl = baseUrl + "/"
	}
	stationItem.Logo = baseUrl + "icon?station_id=" + stationItem.StationId
	stationItem.Relia = "3"
	stationItem.StationLocation = " "
	stationItem.StationBandWidth = " "
	stationItem.StationMime = " "
	stationItem.Bookmark = " "
	return stationItem
}
