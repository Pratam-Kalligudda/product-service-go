package repository

import (
	"github.com/Pratam-Kalligudda/product-service-go/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(domain.Product) error
	UpdateProduct(domain.Product) error
	DeleteProduct(uint) error
	FindProductById(uint) (domain.Product, error)
	FindProductByName(string) (domain.Product, error)
	FindProductByCategory(uint) ([]domain.Product, error)
	FindAllProduct() ([]domain.Product, error)
	// SearchProduct(string) ([]domain.Product, error)
	AddCategory(domain.Category) error
	GetCategories() ([]domain.Category, error)
	GetCategoryByID(uint) (domain.Category, error)
	GetCategoryByName(string) (domain.Category, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product domain.Product) error {
	return r.db.Create(&product).Error
}
func (r *productRepository) UpdateProduct(product domain.Product) error {
	return r.db.Model(&product).Updates(&product).Error
}
func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&domain.Product{}, "id = ?", id).Error
}
func (r *productRepository) FindAllProduct() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r *productRepository) FindProductByCategory(catId uint) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products, "category_id = ?", catId).Error
	return products, err
}
func (r *productRepository) FindProductById(id uint) (domain.Product, error) {
	var product domain.Product
	err := r.db.Find(&product, "id = ?", id).Error
	return product, err
}
func (r *productRepository) FindProductByName(name string) (domain.Product, error) {
	var product domain.Product
	err := r.db.Model(&product).First(&product, "name = ?", name).Error
	return product, err
}
func (r *productRepository) AddCategory(cat domain.Category) error {
	return r.db.Model(&domain.Category{}).Create(&cat).Error
}
func (r *productRepository) GetCategories() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Model(&domain.Category{}).Find(&categories).Error
	return categories, err
}
func (r *productRepository) GetCategoryByID(id uint) (domain.Category, error) {
	var category domain.Category
	err := r.db.Model(domain.Category{}).First(&category, "id = ?", id).Error
	return category, err
}
func (r *productRepository) GetCategoryByName(name string) (domain.Category, error) {
	var category domain.Category
	err := r.db.Model(&domain.Category{}).First(&category, "name = ?", name).Error
	return category, err
}
