package handler

import (
	"github.com/Pratam-Kalligudda/product-service-go/internal/api/rest"
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

func SetupProductHandler(rh rest.HTTPHandler) {
	app := rh.App

	db := repository.NewProductRepository(rh.DB)
	svc := service.NewProductService(db, rh.Helper)
	handler := NewProductHandler(svc)
	pubRoute := app.Group("/products")
	pubRoute.Get("/health", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "service is healthy"})
	})
	pubRoute.Get("/", handler.ListProducts)
	pubRoute.Get("/categories", handler.ListCategories)
	pvtRoute := pubRoute.Group("/", rh.Helper.Authorize)
	pvtRoute.Post("/", handler.AddProduct)
	pvtRoute.Post("/categories", handler.AddCategory)
	pubRoute.Get("/:id", handler.GetProductByID)
	pvtRoute.Put("/:id", handler.UpdateProduct)
	pvtRoute.Delete("/:id", handler.DeleteProduct)

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
