package database

import (
	"fmt"
	"os"

	"github.com/imayrus/url-shortener/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func DatabaseSetup() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	//dbconn := "host=172.17.0.2 user=admin password=test dbname=admin port=5432 sslmode=disable"

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dbdriver := os.Getenv("DB_DRIVER")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s  dbname=%s sslmode=disable", host, port, user,  dbname)

	db, err = gorm.Open(dbdriver, dbinfo)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.ShortUrl{})
	if err != nil {
		fmt.Println(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
