package usercontroller

import (
	"fmt"
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

func StoreUser(c *gin.Context) {

	var request models.UserDataRequest
	errRequest := c.BindJSON(&request)

	if errRequest != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Success: false,
			Message: "Error on your field request",
			Data:    fmt.Sprintf("%v", errRequest),
		})
		return
	}

	createUser := &models.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: request.Password,
	}

	fmt.Print(createUser)

	insert := database.DB.Create(createUser)

	if insert.Error != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Success: false,
			Message: "Error when saving data",
			Data:    fmt.Sprintf("%v", insert.Error),
		})
		return
	} else {
		c.JSON(http.StatusCreated, models.CommonResponse{
			Success: true,
			Message: "Success create new user",
			Data:    createUser,
		})
	}
}
