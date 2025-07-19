package usecase

import "stock-and-order-management/internal/domain"

type orderUsecase struct {
	repo        domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewOrderUsecase(repo domain.OrderRepository, productRepo domain.ProductRepository) domain.OrderUsecase {
	return orderUsecase{
		repo:        repo,
		productRepo: productRepo,
	}
}
