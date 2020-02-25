package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type lock struct {
	status string
}

var vulnLock = lock{status: string("1")}

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Verry vulnerable IoT Lock !")
}

func unLock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unLock called")
	vulnLock.status = "0"
}

func lockTheLock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lockTheLock called")
	vulnLock.status = "1"
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getStatus called")

	_, err := fmt.Fprint(w, vulnLock.status)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
}

func main() {
	fmt.Println("Starting fake IOT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/status/1", lockTheLock).Methods("GET")
	router.HandleFunc("/status/0", unLock).Methods("GET")
	log.Fatal(http.ListenAndServe("192.168.2.225:8080", router))
}
