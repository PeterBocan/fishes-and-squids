package db

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id               string
	Name             string
	Species          string
	OriginCountry    string
	ProductionMethod string
	PricePerTonne    float64
	SupplierName     string
	Certification    string
}

var tablesToUpdate = []any{
	&Product{},
}

// RunMigrations runs through all the tables outlined in the variable above.
func RunMigrations(db *gorm.DB) error {
	// Note this only adds and changes, never removes. That has to be done manually and leaves a lot of mess.
	return db.AutoMigrate(tablesToUpdate...)
}
