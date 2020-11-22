package model

import (
	"ygost/middleware"
)

type StationInfo struct {
	Name    string
	URL     string
	IconURL string `yaml:"iconURL,omitempty"`

	ParentDirName string `yaml:"-"` // internal
	StationId     string `yaml:"-"` // internal
}

type Subdirectory struct {
	Name     string
	Stations []*StationInfo
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

func (si StationInfo) GenerateStationID(subDirName string) string {
	// *** somehow RX-V receivers need the id in a certain format (e.g. MY_41C3D3430E05) ***
	// --> the stationId needs to have exactly 15 characters (longer Ids will result in failing stationlist loads)
	//
	// FIXME for now, we simply hash the station's name plus it's dir name.
	//       then, we prepend "MY_" and finally cut it to 15 chars.
	id := "MY_" + middleware.GenerateStationIDFromName(subDirName+"/"+si.Name)
	id = id[0:15]
	return id
}

var STATIONS MyStationDirectories
var STATIONS_BY_ID map[string]*StationInfo
