package main

import (
	"github.com/keivanipchihagh/shorty/internal"
	"github.com/keivanipchihagh/shorty/internal/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables
	config := configs.NewConfig()

	// Set up logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	internal.Start(config)
}
