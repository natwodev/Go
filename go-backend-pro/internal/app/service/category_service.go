package service

import (
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/repository"
)

type CategoryService interface {
	CreateCategory(name string) (*model.Category, error)
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	UpdateCategory(id uint, name string) (*model.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(name string) (*model.Category, error) {
	category := &model.Category{Name: name}
	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) UpdateCategory(id uint, name string) (*model.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	if err := s.repo.Update(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}
