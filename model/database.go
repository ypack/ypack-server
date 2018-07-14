package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var errorPackageNotFound = errors.New("package %s not found")

// GetPackages retrieve from the database a list of packages from the
// given operating system
func (p *Package) GetPackages(db *gorm.DB) []Package {
	os := p.Versions[0].OS
	var packages []Package

	// search all packages that matches the given OS
	if err := db.Joins("JOIN version on version.os = ? AND version.package_id = package.id", os).
		Find(&packages).Error; err != nil {
		log.Println(err)
		return nil
	}
	// check if found results
	if len(packages) == 0 {
		return []Package{}
	}
	// Get results for each package
	for i := 0; i < len(packages); i++ {
		pkg := &packages[i]
		db.Model(&pkg).Where(&Version{OS: os}).Association("Versions").Find(&pkg.Versions)
		db.Model(&pkg).Where(&Alias{PackageID: pkg.ID}).Association("Alias").Find(&pkg.Alias)
		db.Model(&pkg).Association("Authors").Find(&pkg.Authors)
	}
	// Check for errors
	if err := db.Error; err != nil {
		log.Println(err)
		return nil
	}
	return packages
}

// GetPackageByName retrieve from the database the given package with all versions of the
// given operating system
func (p *Package) GetPackageByName(db *gorm.DB) (Package, error) {
	name, os := p.Name, p.Versions[0].OS
	var pkg Package

	likeName := fmt.Sprintf("%%%s%%", name)
	// search package by name and OS
	if err := db.Joins("JOIN version on version.os = ? AND version.package_id = package.id", os).
		Joins("JOIN alias on alias.package_id = package.id").
		Where("alias.name LIKE ? OR package.name LIKE ?", likeName, likeName).
		Find(&pkg).Error; err != nil {
		// print error and return not found
		log.Println(err)
		return pkg, fmt.Errorf("package %s not found", name)
	}
	// get details for package
	db.Model(&pkg).Where(&Version{OS: os}).Association("Versions").Find(&pkg.Versions)
	db.Model(&pkg).Where(&Alias{PackageID: pkg.ID}).Association("Alias").Find(&pkg.Alias)
	db.Model(&pkg).Association("Authors").Find(&pkg.Authors)

	// Check for errors
	if err := db.Error; err != nil {
		log.Println(err)
		return pkg, fmt.Errorf("unexpected error occurs when getting package %s", name)
	}
	return pkg, nil
}

// GetLatestPackageByName retrieve from the database the latest version of the given package
// for the given operating system
func (p *Package) GetLatestPackageByName(db *gorm.DB) Package {
	log.Printf("Prepared to retrieve from database the latest version of the package %s for %s",
		p.Name, p.Versions[0].OS)
	return Package{}
}
