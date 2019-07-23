package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/saeed4u/go-todo/model"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("There was an error: %s", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword,dbHost,dbPort,dbName)
	log.Printf(connString)
	db_, err := gorm.Open("mysql",connString)
	if err != nil {
		log.Print(err)
	}

	db = db_

	db.Debug().AutoMigrate(&model.User{}, &model.Todo{})
}

func GetDB() *gorm.DB{
	return db
}