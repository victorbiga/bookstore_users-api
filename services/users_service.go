package services

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbiga/bookstore_users-api/domain/users"
)

// CreateUser new user
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil

}

// GetUser get user
func GetUser(c *gin.Context) {

}

// SearchUser search user
func SearchUser(c *gin.Context) {

}
