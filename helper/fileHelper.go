package helper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"ygost/model"
	"ygost/model_yamaha"
)

type Helper interface {
	Unmarshal(in []byte, out interface{}) (err error)
	Marshal(v interface{}) ([]byte, error)
}

type YamlHelper struct {
}

func (YamlHelper) Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}
func (YamlHelper) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

type JSONHelper struct {
}

func (JSONHelper) Unmarshal(in []byte, out interface{}) (err error) {
	return json.Unmarshal(in, out)
}
func (JSONHelper) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

type XMLHelper struct {
}

func (XMLHelper) Unmarshal(in []byte, out interface{}) (err error) {
	return xml.Unmarshal(in, out)
}
func (XMLHelper) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func parse(name string, target interface{}, unmarshaller Helper) interface{} {
	srcFile, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Unknown file: '%s'", name))
	}
	// defer the closing of our srcFile so that we can parse it later on
	defer srcFile.Close()
	// ##################################################################################
	byteValue, err2 := ioutil.ReadAll(srcFile)
	if err2 != nil {
		fmt.Println(err2)
		panic(fmt.Sprintf("Cannot parse file: '%s'", name))
	}
	unmarshaller.Unmarshal(byteValue, target)
	return target
}

func ParseYaml(name string) model.MyStationDirectories {
	var myStations model.MyStationDirectories
	parse(name, &myStations, YamlHelper{})
	return myStations
}
func ParseJSON(name string) model.MyStationDirectories {
	var myStations model.MyStationDirectories
	parse(name, &myStations, JSONHelper{})
	return myStations
}
func ParseYamahaXMLFile(name string) model_yamaha.RootList {
	var listOfItems model_yamaha.RootList
	parse(name, &listOfItems, XMLHelper{})
	return listOfItems
}
