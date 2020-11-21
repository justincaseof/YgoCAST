package model

type StationInfo struct {
	StationName string
	StationURL  string
	IconURL     string // optional
}

type Subdirectory struct {
	Name     string
	Stations []StationInfo
}

type MyStationDirectories struct {
	SubDirectories []Subdirectory `yaml:"my_stations" json:"my_stations"`
}
