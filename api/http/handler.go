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

	url, err := s.UrlService.GetById(int64(id))
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

// GET: /r/:shortened
func (s *HttpApi) Redirect(ctx *gin.Context) {
	shortened := ctx.Param("shortened")

	// Retrieve the URL
	url, err := s.UrlService.GetByShortened(shortened)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the URL is empty or not found, return an error
	if url == nil || url.Original == "" {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url.Original)
}
