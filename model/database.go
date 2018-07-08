package model

import (
	"database/sql"
	"log"
)

// GetPackages retrieve from the database a list of packages from the
// given operating system
func (p *Package) GetPackages(db *sql.DB) []Package {
	log.Printf("Prepared to retrieve from database packages from %s", p.Versions[0].OS)
	return []Package{}
}

// GetPackageByName retrieve from the database the given package with all versions of the
// given operating system
func (p *Package) GetPackageByName(db *sql.DB) []Package {
	log.Printf("Prepared to retrieve from database package %s for %s", p.Name, p.Versions[0].OS)
	return []Package{}
}

// GetLatestPackageByName retrieve from the database the latest version of the given package
// for the given operating system
func (p *Package) GetLatestPackageByName(db *sql.DB) []Package {
	log.Printf("Prepared to retrieve from database the latest version of the package %s for %s",
		p.Name, p.Versions[0].OS)
	return []Package{}
}
