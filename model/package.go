package model

// Package to be installed into a system
type Package struct {
	ID          uint      `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Alias       []Alias   `json:"alias,omitempty"`
	Authors     []Author  `json:"authors" gorm:"many2many:package_author;"`
	Versions    []Version `json:"versions"`
}

// Author of a Package
type Author struct {
	ID      uint   `json:"-"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

// Version of a Package
type Version struct {
	ID        uint   `json:"-"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Checksum  string `json:"checksum,omitempty"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	PackageID uint   `json:"-"`
}

type Alias struct {
	ID        uint   `json:"-"`
	Name      string `json:"name"`
	PackageID uint   `json:"-"`
}
