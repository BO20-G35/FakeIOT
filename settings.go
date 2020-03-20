package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
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

//return true if the xml file is a bomb (not 100% accurate)
func CheckForBomb() bool {

	cmd := exec.Command("/usr/local/bin/python3.7", "/home/tobias/go/src/FakeIOT/xml_read_test.py", userXMLFile)

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if _, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return true
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}
	}
	return false
}
