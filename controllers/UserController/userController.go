package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflisb/backend-service/models"
)

func IndexUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.CommonResponse{
		Success: true,
		Message: "Success get data users",
		Data:    nil,
	})
}
