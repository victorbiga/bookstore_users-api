package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorbiga/bookstore_users-api/domain/users"
	"github.com/victorbiga/bookstore_users-api/services"
	"github.com/victorbiga/bookstore_users-api/utils/errors"
)

// CreateUser new user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
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

// GetUser get user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchUser search user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
