package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/api"
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
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"error": api.ErrNotImplemented.Error()})
}

// DELETE
func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"error": api.ErrNotImplemented.Error()})
}

// GET
func GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
}

// GET
func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"error": api.ErrNotImplemented.Error()})
}
