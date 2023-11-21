package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	// Replace the following values with your MySQL database information
	dbUsername := "root"
	dbPassword := "123Aa@"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "bloggerdb"

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
}

// CloseDB should be deferred after calling ConnectDB
func CloseDB() {
	DB.Close()
}
