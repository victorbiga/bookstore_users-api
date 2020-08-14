package app

import "github.com/victorbiga/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:userid", controllers.GetUser)
	// panic: 'search' in new path '/users/search' conflicts with existing wildcard ':userid' in existing prefix '/users/:userid'
	// router.GET("/users/search", controllers.SearchUser)
}
