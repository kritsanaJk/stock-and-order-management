package repository

import (
	"stock-and-order-management/internal/domain"
	"stock-and-order-management/internal/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (r *orderRepository) Create(order model.Order) error {
	return r.DB.Create(&order).Error
}
