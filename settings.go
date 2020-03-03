package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Settings struct {
	Address string `xml:"address"`
	Port    string `xml:"port"`
}

const defaultXMLFile = "default_config.xml"

func ReadConfigFile() (*Settings, error) {

	data, err := ioutil.ReadFile(defaultXMLFile)
	config := &Settings{}
	_ = xml.Unmarshal([]byte(data), &config)

	return config, err

}
