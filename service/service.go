package service

import (
	"database/sql"
	"github.com/ypack/rest-api-server/model"
	"net/http"
)

// PackageService is the service to retrieve package info
type PackageService struct {
	// DB is the database access layer
	DB *sql.DB
}

// GetPackagesHandler retrieve a list of packages by operating system
func (ps *PackageService) GetPackagesHandler(w http.ResponseWriter, r *http.Request) {
	os := r.URL.Query().Get("os")
	version := model.Version{OS: os}
	pkg := model.Package{Versions: []model.Version{version}}
	// TODO: return packages sorted alphabetically
	pkg.GetPackages(ps.DB)
}

// GetPackageHandler retrieve a package by the given filter
func (ps *PackageService) GetPackageHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//name, os := params.Get("name"), params.Get("os")
}

// GetLatestPackageHandler retrieve the latest version for the given package
func (ps *PackageService) GetLatestPackageHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//name, os := params.Get("name"), params.Get("os")
}
