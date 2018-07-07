package main

import (
	"github.com/gorilla/mux"
	"github.com/ypack/rest-api-server/service"

	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/packages", service.GetPackages).Methods("GET")
	router.HandleFunc("/package", service.GetPackage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
