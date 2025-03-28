package main

import (
	"github.com/keivanipchihagh/shorty/internal"
	"github.com/keivanipchihagh/shorty/internal/configs"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := configs.NewConfig()
	log.SetFormatter(&log.JSONFormatter{})
	internal.Start(config)
}
