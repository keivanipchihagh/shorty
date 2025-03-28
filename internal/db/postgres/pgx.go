package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Option struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	MinConns int
	MaxConns int
}

type PGXDatabase struct {
	*pgxpool.Pool
}

func NewPGXPostgres(option Option) *PGXDatabase {

	config, _ := pgxpool.ParseConfig("")
	config.ConnConfig.Host = option.Host
	config.ConnConfig.Port = uint16(option.Port)
	config.ConnConfig.Database = option.Database
	config.ConnConfig.User = option.Username
	config.ConnConfig.Password = option.Password
	config.MaxConns = int32(option.MaxConns)
	config.MinConns = int32(option.MinConns)

	// Create the connection pool
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		// Log and panic if connection pool creation fails
		log.Panic("unable to create connection pool.")
	}

	// Ping the database to ensure a working connection
	if err := pingConnection(pool); err != nil {
		// Log and panic if the ping test fails
		log.Panic("unable to ping the database.")
	}

	return &PGXDatabase{pool}
}

func pingConnection(pool *pgxpool.Pool) error {
	// Acquire a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("unable to acquire connection: %v", err)
	}
	defer conn.Release() // Ensure the connection is released after use

	// Ping the database to check if the connection is valid
	err = conn.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("failed to ping the database: %v", err)
	}

	// Return nil if the connection is successfully pinged
	return nil
}

func (db *PGXDatabase) Close() {
	// Close the connection pool and releases all resources.
	db.Pool.Close()
}
