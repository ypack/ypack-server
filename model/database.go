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
