package repository

import (
	"stock-and-order-management/internal/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return productRepository{
		DB: db,
	}
}
