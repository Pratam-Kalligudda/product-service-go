package service

import (
	"encoding/json"
	"errors"

	"github.com/Pratam-Kalligudda/product-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/product-service-go/internal/dto"
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
	products, err := s.Repo.FindAllProduct()
	if err != nil {
		return nil, errors.New("error while fetching find all products : " + err.Error())
	}

	if len(products) <= 0 {
		return nil, errors.New("no products found")
	}

	return products, nil
}

func (s *ProductService) GetProductByID(id uint) (domain.Product, error) {
	if id < 0 {
		return domain.Product{}, errors.New("id cannot be negitive")
	}

	product, err := s.Repo.FindProductById(id)
	if err != nil {
		return domain.Product{}, errors.New("error while finding product: " + err.Error())
	}

	if product.ID == 0 {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil
}
func (s *ProductService) GetProductByCategory(catId uint) ([]domain.Product, error) {
	if catId < 0 {
		return nil, errors.New("category id cannot be 0")
	}

	products, err := s.Repo.FindProductByCategory(catId)
	if err != nil {
		return nil, errors.New("error while finidng products by id : " + err.Error())
	}

	if len(products) <= 0 {
		return nil, errors.New("no products of that category id found")
	}

	return products, nil
}
func (s *ProductService) UpdateProduct(id uint, uptProd dto.UpdateProductDTO) (domain.Product, error) {
	if id < 0 {
		return domain.Product{}, errors.New("id cannot be negitive")
	}
	var product domain.Product
	if uptProd.Name != nil {
		product.Name = *uptProd.Name
	}
	if uptProd.Description != nil {
		product.Description = *uptProd.Description
	}
	if uptProd.Stock != nil {
		product.Stock = *uptProd.Stock
	}
	if uptProd.Price != nil {
		product.Price = *uptProd.Price
	}
	if uptProd.CategoryID != nil {
		_, err := s.Repo.GetCategoryByID(*uptProd.CategoryID)
		if err != nil {
			return domain.Product{}, errors.New("no category by that id first create the category")
		}
		product.CategoryID = *uptProd.CategoryID
	}
	product.ID = id
	if err := s.Repo.UpdateProduct(&product); err != nil {
		jsn, _ := json.MarshalIndent(product, "", "\t")
		return domain.Product{}, errors.New("error while updating product :" + err.Error() + " : " + string(jsn))
	}

	product, err := s.Repo.FindProductById(id)
	if err != nil {
		return domain.Product{}, errors.New("error while finding product :" + err.Error())
	}

	return product, nil
}
func (s *ProductService) DeleteProduct(id uint) error {
	if id < 0 {
		return errors.New("id cannot be negative")
	}

	if err := s.Repo.DeleteProduct(id); err != nil {
		return errors.New("error while deleting the product : " + err.Error())
	}

	return nil
}
func (s *ProductService) AddProduct(product domain.Product) (domain.Product, error) {
	_, err := s.Repo.GetCategoryByID(product.CategoryID)
	if err != nil {
		return domain.Product{}, errors.New("no category by that id first create the category")
	}
	_, err = s.Repo.FindProductByName(product.Name)
	if err == nil {
		return domain.Product{}, errors.New("product with the name already exists")
	}

	err = s.Repo.CreateProduct(&product)
	if err != nil {
		return domain.Product{}, errors.New("error while creating product : " + err.Error())
	}

	return product, nil
}
func (s *ProductService) GetCategories() ([]domain.Category, error) {
	categories, err := s.Repo.GetCategories()
	if err != nil {
		return nil, errors.New("error while get categories")
	}
	if len(categories) <= 0 {
		return nil, errors.New("no category created")
	}

	return categories, nil
}
func (s *ProductService) AddCategory(category domain.Category) (domain.Category, error) {
	_, err := s.Repo.GetCategoryByName(category.Name)
	if err == nil {
		return domain.Category{}, errors.New(category.Name + " category already exists")
	}
	err = s.Repo.AddCategory(&category)
	return category, err
}
