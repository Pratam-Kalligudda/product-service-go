package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	PORT   string
	DSN    string
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

	dsn := os.Getenv("DSN")
	if len(dsn) < 0 {
		return ApiConfig{}, errors.New("couldnt load dsn from env")
	}

	secret := os.Getenv("SECRET")
	if len(secret) < 0 {
		return ApiConfig{}, errors.New("couldnt load secret from env")
	}

	return ApiConfig{
		PORT:   port,
		DSN:    dsn,
		SECRET: secret,
	}, nil
}
