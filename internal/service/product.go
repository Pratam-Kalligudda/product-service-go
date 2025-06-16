package service

import (
	"github.com/Pratam-Kalligudda/product-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/product-service-go/internal/helper"
	"github.com/Pratam-Kalligudda/product-service-go/internal/repository"
)

type ProductService struct {
	Repo   repository.ProductRepository
	Helper helper.Helper
}

func NewProductService(repo repository.ProductRepository, helper helper.Helper) *ProductService {
	return &ProductService{Repo: repo, Helper: helper}
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return nil, nil
}
func (s *ProductService) GetProductByID(id uint) (domain.Product, error) {
	return domain.Product{}, nil
}
func (s *ProductService) GetProductByCategory(catId uint) ([]domain.Product, error) {
	return nil, nil
}
func (s *ProductService) UpdateProduct(id uint) error {
	return nil
}
func (s *ProductService) DeleteProduct(id uint) error {
	return nil
}
func (s *ProductService) AddProduct(product any) error {
	return nil
}
func (s *ProductService) GetCategories(u any) ([]domain.Category, error) {
	return nil, nil
}
func (s *ProductService) AddCategory(category any) error {
	return nil
}
