package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "mantap")
}
