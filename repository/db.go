package repository

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("There was an error: %s", err)
	}

	dbHost := os.Getenv("dbHost")
	dbPort := os.Getenv("dbPort")
	dbUsername := os.Getenv("dbUsername")
	dbPassword := os.Getenv("dbPassword")
	dbName := os.Getenv("dbName")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&serverTimezone=UTC", dbUsername, dbPassword,dbHost,dbPort,dbName)

	db_, err := gorm.Open("mysql",connString)
	if err != nil {
		fmt.Print(err)
	}
	db = db_

	db.Debug().AutoMigrate(&model.User{}, &model.Todo{})

	defer db.Close()
}

func GetDB() *gorm.DB{
	return db
}