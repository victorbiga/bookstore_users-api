package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping OK response message
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
