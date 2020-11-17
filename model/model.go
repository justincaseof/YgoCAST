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
	DirCount     string `xml:"DirCount"`

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
