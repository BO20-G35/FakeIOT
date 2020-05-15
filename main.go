package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//kengey
//openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
//openssl genrsa -out server.key 2048

type lock struct {
	status string
}

var vulnLock = lock{status: string("1")}

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Verry vulnerable IoT Lock !")
}

func unLock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unLock called")

	if ValidateKeyForLock(r) {
		vulnLock.status = "0"
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprintf(w, "Lock successfully unlocked.")
	} else {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "Invalid Key.")
	}

}

func lockTheLock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lockTheLock called")

	if ValidateKeyForLock(r) {
		vulnLock.status = "1"
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprintf(w, "Lock successfully locked.")
	} else {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "Invalid Key.")
	}
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getStatus called")

	_, err := fmt.Fprint(w, vulnLock.status)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

}
func getXMLConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getXMLconfig")

	body, _ := ioutil.ReadAll(r.Body)

	// Checks that the body of request was not empty and only continues if it wasn't
	if len(body) <= 0 {
		fmt.Println("Invalid (empty) request.")
		return
	}

	err := SaveXMLFile(body)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "")
	}

	if CheckForBomb() {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, GetFlagString())
		_ = os.Remove(UserXMLFile)

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "Everything ok ;)")
	}

}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Specify path to the python script")
		os.Exit(1)
	}

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
	router.HandleFunc("/status/", getStatus).Methods("GET")
	router.HandleFunc("/status/1", lockTheLock).Methods("GET")
	router.HandleFunc("/status/0", unLock).Methods("GET")
	router.HandleFunc("/config", getXMLConfig).Methods("POST")
	//log.Fatal(http.ListenAndServe(addrString, router))
	log.Fatal(http.ListenAndServeTLS(addrString, "server.crt", "server.key", router))

}
