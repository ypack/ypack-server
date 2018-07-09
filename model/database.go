package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

// GetPackages retrieve from the database a list of packages from the
// given operating system
func (p *Package) GetPackages(db *gorm.DB) []Package {
	os := p.Versions[0].OS
	log.Printf("Prepared to retrieve from database packages from %s", os)

	var packages []Package
	db.Find(&packages)
	log.Println(packages)

	return []Package{}
}

// GetPackageByName retrieve from the database the given package with all versions of the
// given operating system
func (p *Package) GetPackageByName(db *gorm.DB) Package {
	log.Printf("Prepared to retrieve from database package %s for %s", p.Name, p.Versions[0].OS)
	return Package{}
}

// GetLatestPackageByName retrieve from the database the latest version of the given package
// for the given operating system
func (p *Package) GetLatestPackageByName(db *gorm.DB) Package {
	log.Printf("Prepared to retrieve from database the latest version of the package %s for %s",
		p.Name, p.Versions[0].OS)
	return Package{}
}
