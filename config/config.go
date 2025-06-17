package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	PORT   string
	DNS    string
	SECRET string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldnt load env")
		return
	}
}

func SetupConfig() (ApiConfig, error) {
	port := os.Getenv("PORT")
	if len(port) < 0 {
		return ApiConfig{}, errors.New("couldnt load port from env")
	}

	dns := os.Getenv("DNS")
	if len(dns) < 0 {
		return ApiConfig{}, errors.New("couldnt load dsn from env")
	}

	secret := os.Getenv("SECRET")
	if len(secret) < 0 {
		return ApiConfig{}, errors.New("couldnt load secret from env")
	}

	return ApiConfig{
		PORT:   port,
		DNS:    dns,
		SECRET: secret,
	}, nil
}
