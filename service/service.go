package service

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/ypack/rest-api-server/model"
	"github.com/ypack/rest-api-server/validate"
	"log"
	"net/http"
)

// PackageService is the service to retrieve package info
type PackageService struct {
	// DB is the database access layer
	DB *gorm.DB
}

// GetPackagesHandler retrieve a list of packages by operating system
func (ps *PackageService) GetPackagesHandler(w http.ResponseWriter, r *http.Request) {
	os := r.URL.Query().Get("os")
	if err := validate.IsValidOperatingSystem(os); err != nil {
		// return error OS not supported
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	version := model.Version{OS: os}
	pkg := model.Package{Versions: []model.Version{version}}

	packages := pkg.GetPackages(ps.DB)
	if packages == nil {
		// return internal server error
		http.Error(w, "Unexpected error when retrieving packages", http.StatusInternalServerError)
		return
	}
	if len(packages) == 0 {
		// return error 404
		http.Error(w, "No packages found", http.StatusNotFound)
		return
	}
	content, err := json.Marshal(packages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

// GetPackageHandler retrieve a package by the given filter
func (ps *PackageService) GetPackageHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name, os := params.Get("name"), params.Get("os")
	if err := validate.IsValidOperatingSystem(os); err != nil {
		// return error OS not supported
		log.Println(err)
		return
	}
	version := model.Version{OS: os}
	pkg := model.Package{
		Name:     name,
		Versions: []model.Version{version},
	}
	// TODO: return package if found or error if not
	pkg.GetPackageByName(ps.DB)
}

// GetLatestPackageHandler retrieve the latest version for the given package
func (ps *PackageService) GetLatestPackageHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name, os := params.Get("name"), params.Get("os")
	if err := validate.IsValidOperatingSystem(os); err != nil {
		// return error OS not supported
		log.Println(err)
		return
	}
	version := model.Version{OS: os}
	pkg := model.Package{
		Name:     name,
		Versions: []model.Version{version},
	}
	// TODO: return latest version of the package
	pkg.GetLatestPackageByName(ps.DB)
}
