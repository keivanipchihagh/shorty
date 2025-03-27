package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
