package database

import (
	"fmt"

	"github.com/imayrus/url-shortener/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

var db *gorm.DB
var err error

func DatabaseSetup() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	dbconn := "host=172.17.0.2 user=admin password=test dbname=admin port=5432 sslmode=disable"

	host = os.godotenv("DB_HOST")
	user = os.godotenv("")

	db, err = gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.Debug().AutoMigrate(&models.ShortUrl{})
	if err != nil {
		fmt.Println(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
