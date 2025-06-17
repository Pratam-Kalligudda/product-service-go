package api

import (
	"log"

	"github.com/Pratam-Kalligudda/product-service-go/config"
	"github.com/Pratam-Kalligudda/product-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/product-service-go/internal/api/rest/handler"
	"github.com/Pratam-Kalligudda/product-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/product-service-go/internal/helper"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(config config.ApiConfig) {
	app := fiber.New()
	app.Use(logger.New())
	db, err := gorm.Open(postgres.Open(config.DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("error while connecting db")
		return
	}

	log.Print("connection succesfull")
	db.AutoMigrate(&domain.Product{}, &domain.Category{})

	httpHandler := rest.HTTPHandler{
		DB:     db,
		App:    app,
		Helper: helper.Helper{Secret: config.SECRET},
	}
	log.Print("migration succesfull")
	NewHandlerSetup(httpHandler)

	err = app.Listen(config.PORT)
	if err != nil {
		log.Print(err)
	}

}

func NewHandlerSetup(api rest.HTTPHandler) {
	// provide product handler setup method here
	handler.SetupProductHandler(api)
}
