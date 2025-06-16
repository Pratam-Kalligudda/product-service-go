package repository

import (
	"github.com/Pratam-Kalligudda/product-service-go/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(any) error
	UpdateProduct(uint) error
	DeleteProduct(uint) error
	FindProductById(uint) (domain.Product, error)
	FindProductByCategory(uint) ([]domain.Product, error)
	FindAllProduct() ([]domain.Product, error)
	// SearchProduct(string) ([]domain.Product, error)
	AddCategory(any) error
	GetCategories() ([]domain.Category, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product any) error {
	return nil
}
func (r *productRepository) UpdateProduct(id uint) error {
	return nil
}
func (r *productRepository) DeleteProduct(id uint) error {
	return nil
}
func (r *productRepository) FindAllProduct() ([]domain.Product, error) {
	return nil, nil
}
func (r *productRepository) FindProductByCategory(catId uint) ([]domain.Product, error) {
	return nil, nil
}
func (r *productRepository) FindProductById(id uint) (domain.Product, error) {
	return domain.Product{}, nil
}
func (r *productRepository) AddCategory(cat any) error {
	return nil
}
func (r *productRepository) GetCategories() ([]domain.Category, error) {
	return nil, nil
}
