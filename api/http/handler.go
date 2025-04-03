package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/internal/services/urls"
	"github.com/keivanipchihagh/shorty/pkg/models"
	"github.com/sirupsen/logrus"
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

	if !strings.HasPrefix(url.Original, "http://") && !strings.HasPrefix(url.Original, "https://") {
		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}
		url.Original = scheme + "://" + url.Original
	}

	if err := s.UrlService.Create(url); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		logrus.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, url)
}

// GET: /urls/:id
func (s *HttpApi) GetById(ctx *gin.Context) {

	// Parse the 'id' parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalud 'id' parameter"})
		return
	}

	// Retrieve URL information from database
	url, err := s.UrlService.GetById(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		logrus.Error(err)
		return
	}

	// URL not found
	if url == nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.IndentedJSON(http.StatusBadRequest, url)
}

// GET: /urls
func (s *HttpApi) GetAll(ctx *gin.Context) {

	// Retrieve all URLs from the database
	urls, err := s.UrlService.GetAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		logrus.Error(err)
		return
	}

	// No URLs found
	if len(urls) == 0 {
		ctx.Status(http.StatusNoContent)
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
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		logrus.Error(err)
		return
	}

	// If the URL is empty or not found, return an error
	if url == nil || url.Original == "" {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, url.Original)
}
