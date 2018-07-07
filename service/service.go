package service

import (
	"net/http"
)

// GetPackagesHandler retrieve a list of packages by operating system
func GetPackagesHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//os := params["os"]
}

// GetPackageHandler retrieve a package by the given filter
func GetPackageHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//name, os := params.Get("name"), params.Get("os")
}

// GetLatestPackageHandler retrieve the latest version for the given package
func GetLatestPackageHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//name, os := params.Get("name"), params.Get("os")
}
