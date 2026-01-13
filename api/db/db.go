package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NOTE: I am not fond of ORMs, but this will make things easier for starter.

// DB contains all methods which are used for accessing the database (DAO)
type DB struct {
	db *gorm.DB
}

// NewConnection creates the new instance of the DB connection.
func NewConnection(hostname, port, user, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		hostname, user, password, database, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not establish a DB connection: %w", err)
	}
	return conn, nil
}

// NewDB creates the new instance of the DB object.
func NewDB(db *gorm.DB) *DB {
	return &DB{db}
}

// GetAllProducts returns all the products.
func (d *DB) GetAllProducts() ([]Product, error) {
	products := []Product{}
	if err := d.db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("could not find and records")
	}
	return products, nil
}
