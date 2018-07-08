package model

// Package to be installed into a system
type Package struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Website     string  `json:"website"`
	Alias       []Alias `json:"alias"`
	//Authors     []Author  `json:"authors"`
	Versions []Version `json:"versions"`
}

// Author of a Package
type Author struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

// Version of a Package
type Version struct {
	ID        uint   `json:"id"`
	PackageID uint   `json:"package_id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Checksum  string `json:"checksum"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

type Alias struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	PackageID uint   `json:"package_id"`
}
