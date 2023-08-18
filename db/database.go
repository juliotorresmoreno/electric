package db

import (
	"log"

	"github.com/juliotorresmoreno/electric/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil {
		cxn := ":memory:?cache=shared"
		connection, err = gorm.Open(sqlite.Open(cxn), &gorm.Config{})
	}

	return connection, err
}

func AutoMigrate(db *gorm.DB) {
	conn, err := GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	conn.AutoMigrate(&models.Country{})
	conn.AutoMigrate(&models.City{})
	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Location{})
	conn.AutoMigrate(&models.Consumption{})
}
