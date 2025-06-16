package handler

import (
	"github.com/Pratam-Kalligudda/product-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/product-service-go/internal/helper"
	"github.com/Pratam-Kalligudda/product-service-go/internal/repository"
	"github.com/Pratam-Kalligudda/product-service-go/internal/service"
	"github.com/gofiber/fiber/v3"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) ProductHandler {
	return ProductHandler{service: service}
}

func SetupProductHandler(api rest.HTTPHandler) {
	app := api.App

	db := repository.NewProductRepository(api.DB)
	svc := service.NewProductService(db, helper.Helper{Secret: api.Config.SECRET})
	handler := NewProductHandler(svc)

	app.Get("/products", handler.ListProducts)
	app.Get("/products/:id", handler.GetProductByID)
	app.Post("/products", handler.AddProduct)
	app.Put("/products/:id", handler.UpdateProduct)
	app.Delete("/products/:id", handler.DeleteProduct)
	app.Get("/categories", handler.ListCategories)
	app.Post("/categories", handler.AddCategory)

	// app.Listen(api.Config.PORT)
}

func (h *ProductHandler) ListProducts(ctx fiber.Ctx) error {
	return nil
}

func (h *ProductHandler) GetProductByID(ctx fiber.Ctx) error {
	return nil
}
func (h *ProductHandler) AddProduct(ctx fiber.Ctx) error {
	return nil
}
func (h *ProductHandler) UpdateProduct(ctx fiber.Ctx) error {
	return nil
}
func (h *ProductHandler) DeleteProduct(ctx fiber.Ctx) error {
	return nil
}
func (h *ProductHandler) ListCategories(ctx fiber.Ctx) error {
	return nil
}
func (h *ProductHandler) AddCategory(ctx fiber.Ctx) error {
	return nil
}
