package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func GetConfig() *Config {
	godotenv.Load(".env.production", ".env", ".env.development")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port: port,
	}
}
