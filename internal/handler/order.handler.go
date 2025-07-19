package handler

import (
	"stock-and-order-management/internal/domain"
)

type orderHandler struct {
	usecase domain.OrderUsecase
}

func NewOrderHandler(usecase domain.OrderRepository) *orderHandler {
	return &orderHandler{
		usecase: usecase,
	}
}
