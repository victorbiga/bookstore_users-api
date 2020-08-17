package services

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbiga/bookstore_users-api/domain/users"
	"github.com/victorbiga/bookstore_users-api/utils/errors"
)

// CreateUser new user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser get user
func GetUser(c *gin.Context) {

}

// SearchUser search user
func SearchUser(c *gin.Context) {

}
