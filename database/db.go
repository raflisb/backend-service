package database

import (
	"fmt"
	"os"

	"github.com/raflisb/backend-service/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		message := fmt.Sprintf("%d", err)
		fmt.Println(message)
	}

	DB = db
	fmt.Println("Connected to database ... ")
	migration()
}

func migration() {
	// Register the model to migrate here
	DB.AutoMigrate(&model.User{})
}
