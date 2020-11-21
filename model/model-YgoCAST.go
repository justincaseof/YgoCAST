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
	SubDirectories map[string]Subdirectory `yaml:"my_stations" json:"my_stations"`
}

func (msd MyStationDirectories) SubDirectoriesAsList() []Subdirectory {
	result := make([]Subdirectory, 0, len(msd.SubDirectories))
	for _, v := range msd.SubDirectories {
		result = append(result, v)
	}
	return result
}
