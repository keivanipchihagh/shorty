package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/api"
	"github.com/keivanipchihagh/shorty/internal/services/urls"
	"github.com/keivanipchihagh/shorty/pkg/models"
)

type HttpApi struct {
	UrlService urls.UrlService
}

func NewHttpApi(UrlService urls.UrlService) *HttpApi {
	return &HttpApi{UrlService: UrlService}
}

// POST
func (s *HttpApi) Create(ctx *gin.Context) {
	var url *models.URL
	if err := ctx.BindJSON(&url); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	url, err := s.UrlService.Create(url)
	fmt.Println(url, err)
}

// GET
func (s *HttpApi) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "valid 'id' is required"})
		return
	}

	url, err := s.UrlService.GetById(id)
	fmt.Println(url, err)
}

// GET
func (s *HttpApi) GetAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusNotImplemented, gin.H{"error": api.ErrNotImplemented.Error()})

	urls, err := s.UrlService.GetAll()
	fmt.Println(urls, err)
}
