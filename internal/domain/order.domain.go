package domain

import (
	"stock-and-order-management/internal/dto"
	"stock-and-order-management/internal/model"
)

type OrderUsecase interface {
	Create(req dto.CreateOrderRequest) error
}

type OrderRepository interface {
	Create(product model.Order) error
}
