package handler

import "stock-and-order-management/internal/domain"

type productHandler struct {
	usecase domain.ProductUsecase
}

func NewProducthandler(usecase domain.ProductRepository) *productHandler {
	return &productHandler{
		usecase: usecase,
	}
}
