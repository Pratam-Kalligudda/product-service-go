package rest

import (
	"github.com/Pratam-Kalligudda/product-service-go/config"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type HTTPHandler struct {
	DB     *gorm.DB
	App    *fiber.App
	Config config.ApiConfig
}
