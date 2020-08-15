package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorbiga/bookstore_users-api/domain/users"
	"github.com/victorbiga/bookstore_users-api/utils/errors"
)

// CreateUser new user
func CreateUser(user users.User) (users.User, *errors.RestErr) {
	return user, nil

	var defaultUser users.User

	return defaultUser, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}

}

// GetUser get user
func GetUser(c *gin.Context) {

}

// SearchUser search user
func SearchUser(c *gin.Context) {

}
