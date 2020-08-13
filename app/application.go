package app

import ( "github.com/gin-gonic/gin")

var (
	router = gin.Default()
) type

func StartApplication() {
	mapUrls()
	router.Run( addr: ":8080")

}