package model

import (
	"database/sql"
	"log"
)

// GetPackages retrieve from the database a list of packages from the
// given operating system
func (p *Package) GetPackages(db *sql.DB) []Package {
	os := p.Versions[0].OS
	log.Printf("Prepared to retrieve from database packages from %s", os)

	stmt, err := db.Prepare(queryGetPackagesByOS)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query(os)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	// Transform data to the know model Package
	var packages []Package
	for rows.Next() {
		var pkg Package
		// FIXME: expected 11 params
		err = rows.Scan(&pkg.Name, &pkg.Description, &pkg.Website)
		if err != nil {
			log.Println(err)
			break
		}
		packages = append(packages, pkg)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil
	}
	return packages
}

// GetPackageByName retrieve from the database the given package with all versions of the
// given operating system
func (p *Package) GetPackageByName(db *sql.DB) Package {
	log.Printf("Prepared to retrieve from database package %s for %s", p.Name, p.Versions[0].OS)
	return Package{}
}

// GetLatestPackageByName retrieve from the database the latest version of the given package
// for the given operating system
func (p *Package) GetLatestPackageByName(db *sql.DB) Package {
	log.Printf("Prepared to retrieve from database the latest version of the package %s for %s",
		p.Name, p.Versions[0].OS)
	return Package{}
}
