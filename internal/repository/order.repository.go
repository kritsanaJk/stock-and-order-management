package repository

import (
	"stock-and-order-management/internal/domain"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return orderRepository{
		DB: db,
	}
}
