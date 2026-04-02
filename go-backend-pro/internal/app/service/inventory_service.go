package service

import (
	"errors"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/repository"
	"gorm.io/gorm"
)

type InventoryService interface {
	CreateInventory(productID uint, quantity int, location string) (*model.Inventory, error)
	GetAllInventories() ([]model.Inventory, error)
	GetInventoryByID(id uint) (*model.Inventory, error)
	IncreaseInventory(id uint, amount int) (*model.Inventory, error)
	DecreaseInventory(id uint, amount int) (*model.Inventory, error)
	DeleteInventory(id uint) error
}

type inventoryService struct {
	repo        repository.InventoryRepository
	productRepo repository.ProductRepository
}

var ErrInventoryInsufficientQuantity = errors.New("insufficient quantity")

func NewInventoryService(
	repo repository.InventoryRepository,
	productRepo repository.ProductRepository,
) InventoryService {
	return &inventoryService{repo: repo, productRepo: productRepo}
}

func (s *inventoryService) CreateInventory(productID uint, quantity int, location string) (*model.Inventory, error) {
	if _, err := s.productRepo.FindByID(productID); err != nil {
		return nil, err
	}
	inventory := &model.Inventory{
		ProductID: productID,
		Quantity:  quantity,
		Location:  location,
	}
	if err := s.repo.Create(inventory); err != nil {
		return nil, err
	}
	return inventory, nil
}

func (s *inventoryService) GetAllInventories() ([]model.Inventory, error) {
	return s.repo.FindAll()
}

func (s *inventoryService) GetInventoryByID(id uint) (*model.Inventory, error) {
	return s.repo.FindByID(id)
}

func (s *inventoryService) IncreaseInventory(id uint, amount int) (*model.Inventory, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	return s.repo.Increase(id, amount)
}

func (s *inventoryService) DecreaseInventory(id uint, amount int) (*model.Inventory, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	inventory, err := s.repo.Decrease(id, amount)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInventoryInsufficientQuantity
	}
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (s *inventoryService) DeleteInventory(id uint) error {
	return s.repo.Delete(id)
}
