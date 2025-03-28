package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/internal/services/urls"
	"github.com/keivanipchihagh/shorty/pkg/models"
)

type HttpApi struct {
	UrlService urls.UrlService
}

func NewHttpApi(UrlService urls.UrlService) *HttpApi {
	return &HttpApi{UrlService: UrlService}
}

// POST: /urls
func (s *HttpApi) Create(ctx *gin.Context) {
	var url *models.URL
	if err := ctx.BindJSON(&url); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := s.UrlService.Create(url); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, url)
}

// GET: /urls/:id
func (s *HttpApi) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "valid 'id' is required"})
		return
	}

	url, err := s.UrlService.GetById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusBadRequest, url)
}

// GET: /urls
func (s *HttpApi) GetAll(ctx *gin.Context) {
	urls, err := s.UrlService.GetAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, urls)
}
