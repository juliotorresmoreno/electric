package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/juliotorresmoreno/electric/config"
	"github.com/juliotorresmoreno/electric/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connection *gorm.DB

func GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil {
		conf, _ := config.GetConfig()
		dbConf := conf.Database
		fmt.Println(dbConf)
		switch dbConf["driver"] {
		case "postgres":
			connection, err = makePostgreSQLConnection(dbConf)
		case "sqlite":
			fallthrough
		default:
			connection, err = makeSQLiteConnection(dbConf)
		}
	}

	return connection, err
}

func makeSQLiteConnection(dbConf map[string]string) (*gorm.DB, error) {
	path := dbConf["path"]
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	return db, err
}

func makePostgreSQLConnection(dbConf map[string]string) (*gorm.DB, error) {
	dsn := dbConf["dsn"]
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
	db.Logger = newLogger

	return db, err
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
	conn.AutoMigrate(&models.Metric{})
	conn.AutoMigrate(&models.Consumption{})

	count := int64(0)
	db.Model(models.Country{}).Count(&count)
	fmt.Println(count)
	if count > 0 {
		return
	}
	country := &models.Country{
		Name: "Colombia",
		Code: "CO",
	}
	conn.Create(country)
	city := &models.City{
		Name:      "Bogota",
		Code:      "BOG",
		CountryId: country.ID,
	}
	conn.Create(city)
	user := &models.User{
		FirstName: "Jhon",
		LastName:  "Doe",
		Phone:     "+5 555 5555",
	}
	conn.Create(user)
	location := &models.Location{
		OwnerId:     user.ID,
		Address:     "None",
		Description: "",
		Latitude:    132.24,
		Longitude:   251.25,
	}
	conn.Create(location)
	metric := &models.Metric{
		LocationId:         location.ID,
		Date:               time.Now(),
		Active:             100,
		ReactiveInductive:  10,
		ReactiveCapacitive: 120,
		Exported:           30,
	}
	conn.Create(metric)
}
