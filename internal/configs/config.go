package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Http     HttpConfig
	Postgres PostgresConfig
}

type HttpConfig struct {
	Host string
	Port int
}

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	MaxConns int
	MinConns int
}

func NewConfig() *Config {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Panic("error loading .env")
	}
	viper.AutomaticEnv()

	http := HttpConfig{
		Host: viper.GetString("HTTP_HOST"),
		Port: viper.GetInt("HTTP_PORT"),
	}

	postgres := PostgresConfig{
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetInt("POSTGRES_PORT"),
		Username: viper.GetString("POSTGRES_USERNAME"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		Database: viper.GetString("POSTGRES_DATABASE"),
		MaxConns: viper.GetInt("POSTGRES_MAX_CONNS"),
		MinConns: viper.GetInt("POSTGRES_MIN_CONNS"),
	}

	return &Config{
		Http:     http,
		Postgres: postgres,
	}
}
