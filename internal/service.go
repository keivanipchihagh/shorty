package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/api/http"
	"github.com/keivanipchihagh/shorty/internal/configs"
	"github.com/keivanipchihagh/shorty/internal/db/postgres"
	"github.com/keivanipchihagh/shorty/internal/metrics"
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

	router := gin.Default()
	// Register Middlewares
	router.Use(metrics.PrometheusMetrics())
	// Register routes
	router.POST("/urls", http.Create)
	router.GET("/urls", http.GetAll)
	router.GET("/urls/:id", http.GetById)
	router.PUT("/urls/:id", http.Update)
	router.DELETE("/urls/:id", http.Delete)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	address := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)
	router.Run(address)
}
