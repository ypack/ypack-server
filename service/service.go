package service

import (
	"net/http"
)

// GetPackages retrieve a list of packages by operating system
func GetPackages(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//os := params["os"]
}

// GetPackage retrieve a package by the given filter
func GetPackage(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//name, os := params.Get("name"), params.Get("os")
}
