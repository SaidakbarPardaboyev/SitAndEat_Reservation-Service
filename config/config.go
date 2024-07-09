package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	RESERVATION_SERVICE string
	DB_HOST             string
	DB_PORT             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
}

func Load() Config {
	if err := godotenv.Load("/home/azizbek/github/Restaurant/.env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "hamidjon4424"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "resuserservice"))
	config.RESERVATION_SERVICE = cast.ToString(Coalesce("RESERVATION_SERVICE", "50051"))
	config.RESERVATION_SERVICE = cast.ToString(Coalesce("RESERVATION_SERVICE", "6666"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
