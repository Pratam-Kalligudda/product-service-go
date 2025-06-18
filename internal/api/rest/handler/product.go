package handler

import (
	"strconv"

	"github.com/Pratam-Kalligudda/product-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/product-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/product-service-go/internal/dto"
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

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "service is healthy"})
	})
	//public endpoints
	pubRoute := app.Group("/product")
	pubRoute.Get("/", handler.ListProducts)
	pubRoute.Get("/category", handler.ListCategories)
	pubRoute.Get("/:id", handler.GetProductByID)

	//private endpoints
	pvtRoute := pubRoute.Group("/", rh.Helper.Authorize)
	pvtRoute.Post("/", handler.AddProduct)
	pvtRoute.Post("/category", handler.AddCategory)
	pvtRoute.Put("/:id", handler.UpdateProduct)
	pvtRoute.Delete("/:id", handler.DeleteProduct)

	// app.Listen(api.Config.PORT)
}

func (h *ProductHandler) ListProducts(ctx fiber.Ctx) error {
	products, err := h.service.GetProducts()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (h *ProductHandler) GetProductByID(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	if idStr == "" || len(idStr) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "incorrect params"})
	}
	prodID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "format error"})
	}
	product, err := h.service.GetProductByID(uint(prodID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "found", "product": product})
}

func (h *ProductHandler) AddProduct(ctx fiber.Ctx) error {
	var product domain.Product
	err := ctx.Bind().Body(&product)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while binding :" + err.Error()})
	}
	// userId := ctx.Locals("userId", 0)
	// if userId == 0 {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error with userID"})
	// }
	product, err = h.service.AddProduct(product)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while adding :" + err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "product added succesfully :",
		"product": product,
	})
}

func (h *ProductHandler) UpdateProduct(ctx fiber.Ctx) error {
	var productDto dto.UpdateProductDTO
	err := ctx.Bind().Body(&productDto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while binding :" + err.Error()})
	}

	idStr := ctx.Params("id")
	if idStr == "" || len(idStr) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "incorrect params"})
	}

	prodID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "format error"})
	}

	product, err := h.service.UpdateProduct(uint(prodID), productDto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while updating :" + err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "product updated succesfully :",
		"product": product,
	})
}

func (h *ProductHandler) DeleteProduct(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	if idStr == "" || len(idStr) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "incorrect params"})
	}

	prodID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "format error"})
	}
	err = h.service.DeleteProduct(uint(prodID))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while deletin:" + err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "product deleted succesfully :",
	})
}

func (h *ProductHandler) ListCategories(ctx fiber.Ctx) error {
	categories, err := h.service.GetCategories()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(categories)
}

func (h *ProductHandler) AddCategory(ctx fiber.Ctx) error {
	var category domain.Category
	err := ctx.Bind().Body(&category)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while binding :" + err.Error()})
	}
	// userId := ctx.Locals("userId", 0)
	// if userId == 0 {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error with userID"})
	// }

	category, err = h.service.AddCategory(category)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error while adding :" + err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "category added succesfully :",
		"category": category,
	})
}
