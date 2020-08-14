package app

import (
	"github.com/victorbiga/bookstore_users-api/controllers/ping"
	"github.com/victorbiga/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:userid", users.GetUser)
	// panic: 'search' in new path '/users/search' conflicts with existing wildcard ':userid' in existing prefix '/users/:userid'
	// router.GET("/users/search", controllers.SearchUser)
}
