package service

import (
	"log"
	"net/http"
)

// GetPackages retrieve a list of packages by operating system
func GetPackages(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling get packages")
}

// GetPackage retrieve a package by the given filter
func GetPackage(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling get package")
}
