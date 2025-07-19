package domain

import (
	"stock-and-order-management/internal/dto"
	"stock-and-order-management/internal/model"
)

type ProductUsecase interface {
	Create(req dto.CreateProductRequest) error
	List(res *[]dto.ProductResponse) error
	Product(productID int, res *dto.ProductResponse) error
}

type ProductRepository interface {
	Create(product model.Product) error
	CheckDuplicate(name string, product *model.Product) error
	List(res *[]model.Product) error
	GetByID(productID int, res *model.Product) error
	Update(productID int, req model.Product) error
}
