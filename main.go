package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

// her må vi ha sikkerhetshullet
func getXMLConfig(w http.ResponseWriter, r *http.Request) {
	//her er da XML data som blir sendt via POST
	body, _ := ioutil.ReadAll(r.Body)

	//TODO lagre til fil og restart http listener
	fmt.Println(body)

}

func main() {
	fmt.Println("Starting fake IOT")

	config, err := ReadConfigFile()
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(config.Address)
	fmt.Println(config.Port)

	addrString := config.Address + ":" + config.Port

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/status/1", lockTheLock).Methods("GET")
	router.HandleFunc("/status/0", unLock).Methods("GET")
	router.HandleFunc("/config", getXMLConfig).Methods("GET", "POST")
	log.Fatal(http.ListenAndServe(addrString, router))

}
