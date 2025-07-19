package usecase

import "stock-and-order-management/internal/domain"

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return productUsecase{
		repo: repo,
	}
}
