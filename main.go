package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"ygost/model"
)

func main() {
	fmt.Println("Hi!")
}

func ParseFile(name string) {
	// ##################################################################################
	xmlFile, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened file " + name)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// ##################################################################################
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var listOfItems model.ListOfItems
	xml.Unmarshal(byteValue, &listOfItems)

	fmt.Println("done.", listOfItems)
}
