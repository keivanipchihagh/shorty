package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/pkg/models"
)

// POST
func Create(c *gin.Context) {
	var url models.URL
	if err := c.BindJSON(&url); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
}

// PUT
func Update(c *gin.Context) {
	var url models.URL
	if err := c.BindJSON(&url); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
}

// DELETE
func Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
}

// GET
func Get(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
}

// GET
func GetAll(c *gin.Context) {

}
