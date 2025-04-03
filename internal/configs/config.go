package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Http     HttpConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
	TTL      int
}

type HttpConfig struct {
	Host string
	Port int
	Mode string
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
		Mode: viper.GetString("HTTP_MODE"),
	}

	redis := RedisConfig{
		Host:     viper.GetString("REDIS_HOST"),
		Port:     viper.GetInt("REDIS_PORT"),
		Password: viper.GetString("REDIS_PASSWORD"),
		Database: viper.GetInt("REDIS_DATABASE"),
		TTL:      viper.GetInt("REDIS_TTL"),
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
		Redis:    redis,
	}
}
