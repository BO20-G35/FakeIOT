package main

import (
	"io/ioutil"
	"net/http"
)

const passwordFile = "lockPassword.txt"
const flaggFile = "flag.txt"

func ValidateKeyForLock(r *http.Request) bool {

	key := r.URL.Query().Get("k")

	rightKey, _ := ioutil.ReadFile(passwordFile)

	return key == string(rightKey)
}

func GetFlagString() string {
	flag, _ := ioutil.ReadFile(flaggFile)
	return string(flag)
}
