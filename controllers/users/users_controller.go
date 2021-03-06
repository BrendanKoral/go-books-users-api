package users

import (
	"github.com/brendankoral/go-books-users-api/domain/users"
	"github.com/brendankoral/go-books-users-api/services"
	"github.com/brendankoral/go-books-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
