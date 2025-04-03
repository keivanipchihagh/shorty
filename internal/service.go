package internal

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/api/http"
	"github.com/keivanipchihagh/shorty/internal/configs"
	"github.com/keivanipchihagh/shorty/internal/db/postgres"
	"github.com/keivanipchihagh/shorty/internal/metrics"
	"github.com/keivanipchihagh/shorty/internal/services/urls"
	"github.com/keivanipchihagh/shorty/pkg/repositories"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(config *configs.Config) {

	db := postgres.NewPGXPostgres(postgres.Option{
		Host:     config.Postgres.Host,
		Port:     config.Postgres.Port,
		Username: config.Postgres.Username,
		Password: config.Postgres.Password,
		Database: config.Postgres.Database,
		MinConns: config.Postgres.MinConns,
		MaxConns: config.Postgres.MaxConns,
	})
	defer db.Close()

	urlRepo := repositories.NewUrlRepo(db.Pool)
	urlService := urls.NewUrlService(urlRepo)
	httpApi := http.NewHttpApi(urlService)

	gin.SetMode(config.Http.Mode)
	router := gin.Default()
	log.Info(fmt.Sprintf("Gin running in %v mode", config.Http.Mode))

	// Register middlewares
	router.Use(metrics.PrometheusMetrics())
	// Register routes
	router.GET("/:shortened", httpApi.Redirect)
	router.POST("/urls", httpApi.Create)
	router.GET("/urls", httpApi.GetAll)
	router.GET("/urls/:id", httpApi.GetById)
	router.GET("/health", httpApi.Health)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	address := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)
	log.Info(fmt.Sprintf("Gin running on %v", address))
	if err := router.Run(address); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
