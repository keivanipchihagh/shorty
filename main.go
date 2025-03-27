package main

import (
	"github.com/keivanipchihagh/shorty/internal"
	"github.com/keivanipchihagh/shorty/internal/db/postgres"
)

func main() {

	option := postgres.Option{
		Host:           "localhost",
		Port:           5432,
		Username:       "postgres",
		Password:       "postgres",
		Database:       "postgres",
		MinConnections: 1,
		MaxConnections: 3,
	}
	db := postgres.NewPGXPostgres(option)
	defer db.Close()

	internal.Start()
}
