package helper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"ygost/model"
)

func ParseFile(name string) model.ListOfItems {
	// ##################################################################################
	xmlFile, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// ##################################################################################
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var listOfItems model.ListOfItems
	xml.Unmarshal(byteValue, &listOfItems)

	return listOfItems
}
