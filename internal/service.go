package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/shorty/api/http"
	"github.com/keivanipchihagh/shorty/internal/configs"
	"github.com/keivanipchihagh/shorty/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() {

	configs := configs.NewConfig()
	router := gin.Default()

	// Register Prometheus middleware
	router.Use(metrics.PrometheusMetrics())

	// Register API routes
	router.GET("/", http.HelloWorld)

	// Expose Prometheus metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	address := fmt.Sprintf("%s:%d", configs.Host, configs.Port)
	router.Run(address)
}
