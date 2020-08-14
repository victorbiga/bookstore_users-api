package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorbiga/bookstore_users-api/domain/users"
	"github.com/victorbiga/bookstore_users-api/services"
)

// CreateUser new user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO: handle json error
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Hamdle json error
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
