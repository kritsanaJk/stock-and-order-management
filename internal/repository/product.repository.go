package repository

import (
	"stock-and-order-management/internal/domain"
	"stock-and-order-management/internal/model"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (r *productRepository) Create(product model.Product) error {
	return r.DB.Create(&product).Error
}

func (r *productRepository) CheckDuplicate(name string, product *model.Product) error {
	return r.DB.Debug().Model(&model.Product{}).Where("lower(name) = lower(?)", name).First(&product).Error
}
