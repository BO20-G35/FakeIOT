package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Settings struct {
	Address string `xml:"address"`
	Port    string `xml:"port"`
}

const defaultXMLFile = "default_config.xml"
const userXMLFile = "user_config.xml"

func ReadConfigFile() (*Settings, error) {

	var fileToRead = userXMLFile
	if _, err := os.Stat(userXMLFile); err != nil {
		fileToRead = defaultXMLFile
		fmt.Println("userXMLFile not found reading from defaultXMLFile")
	}

	fmt.Println("reading userXMLFIle")
	data, err := ioutil.ReadFile(fileToRead)
	config := &Settings{}
	_ = xml.Unmarshal([]byte(data), &config)

	return config, err

}

func SaveXMLFile(data []byte) error {

	fmt.Println("Saving new userXMLFile")
	err := ioutil.WriteFile(userXMLFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
