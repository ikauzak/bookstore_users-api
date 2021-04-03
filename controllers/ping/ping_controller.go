package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Ping responds if the server is ok
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
