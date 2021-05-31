package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	usercontroller "github.com/raflisb/backend-service/controllers/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.GET("/user", usercontroller.IndexUsers)
	r.POST("/user", usercontroller.StoreUser)
	return r
}
