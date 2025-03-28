package main

import (
	"github.com/keivanipchihagh/shorty/internal"
	"github.com/keivanipchihagh/shorty/internal/configs"
)

func main() {
	config := configs.NewConfig()
	internal.Start(config)
}
