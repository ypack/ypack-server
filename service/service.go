package service

import (
	"log"
	"net/http"
)

// GetPackages retrieve a list of packages by operating system
func GetPackages(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	log.Println("Calling get packages for", params["os"])
}

// GetPackage retrieve a package by the given filter
func GetPackage(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	log.Printf("Calling get package %s for %s\n", params.Get("name"), params.Get("os"))
}
