package model

// Package to be installed into a system
type Package struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Alias       []string  `json:"alias"`
	Authors     []Author  `json:"authors"`
	Versions    []Version `json:"versions"`
}

// Author of a Package
type Author struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

// Version of a Package
type Version struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Checksum string `json:"checksum"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
}
