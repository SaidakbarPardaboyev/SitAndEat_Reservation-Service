package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT   string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "hamidjon4424"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "resuserservice"))
	config.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", "50051"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
