package main

import (
	"github.com/gorilla/mux"
	"github.com/ypack/rest-api-server/logger"
	"github.com/ypack/rest-api-server/service"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Use(logger.Logger)

	v1group := router.PathPrefix("/v1").Subrouter()
	v1group.HandleFunc("/packages", service.GetPackages).Methods("GET")
	v1group.HandleFunc("/package", service.GetPackage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
