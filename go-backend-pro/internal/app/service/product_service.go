package service

import (
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/repository"
)

type ProductService interface {
	CreateProduct(name, description string, price float64, categoryID uint) (*model.Product, error)
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(id uint, name, description string, price float64, categoryID uint) (*model.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(name, description string, price float64, categoryID uint) (*model.Product, error) {
	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}
	if err := s.repo.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetAllProducts() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(id uint, name, description string, price float64, categoryID uint) (*model.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	product.Name = name
	product.Description = description
	product.Price = price
	product.CategoryID = categoryID
	if err := s.repo.Update(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
