package model

import "ygost/middleware"

type StationInfo struct {
	StationName   string
	StationURL    string
	IconURL       string `yaml:"iconURL,omitempty"`
	ParentDirName string `yaml:"-"`
}

type Subdirectory struct {
	Name     string
	Stations []StationInfo
}

type MyStationDirectories struct {
	SubDirectories map[string]Subdirectory `yaml:"my_stations" json:"my_stations"`
}

func (msd MyStationDirectories) SubDirectoriesAsList() []Subdirectory {
	result := make([]Subdirectory, 0, len(msd.SubDirectories))
	for _, v := range msd.SubDirectories {
		result = append(result, v)
	}
	return result
}

func (si *StationInfo) GenerateStationID(subDirName string) string {
	// FIXME for now, we simply hash the station's name.
	return "MY_"+middleware.CalculateStationID(subDirName + "/" + si.StationName)
}

var STATIONS MyStationDirectories
var STATIONS_BY_ID map[string]*StationInfo
