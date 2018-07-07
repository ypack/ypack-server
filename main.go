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
	v1group.HandleFunc("/packages", service.GetPackagesHandler).Methods("GET")
	v1group.HandleFunc("/package", service.GetPackageHandler).Methods("GET")
	v1group.HandleFunc("/package/latest", service.GetLatestPackageHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
