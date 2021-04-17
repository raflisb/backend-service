package main

import ( 
	// "net/http"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	// "github.com/joho/godotenv'
	"fmt"
	"github.com/joho/godotenv"
	"github.com/raflisb/backend-service/database"
	"github.com/raflisb/backend-service/model"

)


func main() { 
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	database.Init() 

	DB:= database.DB
	
	DB.AutoMigrate(&model.User{})
}