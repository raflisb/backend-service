package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflisb/backend-service/database"
	"github.com/raflisb/backend-service/models"
	"github.com/raflisb/backend-service/pagination"
)

func IndexUsers(c *gin.Context) {
	var paginationQuery models.PaginationRequest
	c.BindQuery(&paginationQuery)

	var users []models.User

	paginator := pagination.Paging(&pagination.Param{
		DB:      database.DB,
		Page:    paginationQuery.Page,
		Limit:   paginationQuery.PerPage,
		OrderBy: []string{"id desc"},
		ShowSQL: false,
	}, &users)

	c.JSON(http.StatusOK, paginator)
}
