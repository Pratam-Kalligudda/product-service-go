package api

import (
	"log"

	"github.com/Pratam-Kalligudda/product-service-go/config"
	"github.com/Pratam-Kalligudda/product-service-go/internal/api/rest"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(config config.ApiConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error while connecting db")
		return
	}

	log.Print("connection succesfull")
	db.AutoMigrate()

	httpHandler := rest.HTTPHandler{
		DB:  db,
		App: app,
	}

	NewHandlerSetup(httpHandler)

	app.Listen(config.PORT)

}

func NewHandlerSetup(handler rest.HTTPHandler) {
	// provide product handler setup method here
}
