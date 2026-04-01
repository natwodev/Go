package repository

import (
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Create(inventory *model.Inventory) error
	FindAll() ([]model.Inventory, error)
	FindByID(id uint) (*model.Inventory, error)
	Increase(id uint, amount int) (*model.Inventory, error)
	Decrease(id uint, amount int) (*model.Inventory, error)
	Delete(id uint) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) Create(inventory *model.Inventory) error {
	return r.db.Create(inventory).Error
}

func (r *inventoryRepository) FindAll() ([]model.Inventory, error) {
	var inventories []model.Inventory
	err := r.db.Preload("Product").Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) FindByID(id uint) (*model.Inventory, error) {
	var inventory model.Inventory
	err := r.db.Preload("Product").First(&inventory, id).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepository) Increase(id uint, amount int) (*model.Inventory, error) {
	if err := r.db.Model(&model.Inventory{}).
		Where("id = ?", id).
		Update("quantity", gorm.Expr("quantity + ?", amount)).Error; err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *inventoryRepository) Decrease(id uint, amount int) (*model.Inventory, error) {
	result := r.db.Model(&model.Inventory{}).
		Where("id = ? AND quantity >= ?", id, amount).
		Update("quantity", gorm.Expr("quantity - ?", amount))
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return r.FindByID(id)
}

func (r *inventoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Inventory{}, id).Error
}
