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
	return r.DB.Model(&model.Product{}).Where("lower(name) = lower(?)", name).First(&product).Error
}

func (r *productRepository) List(res *[]model.Product) error {
	return r.DB.Model(&model.Product{}).Order("id asc").Find(&res).Error
}

func (r *productRepository) GetByID(productID int, res *model.Product) error {
	return r.DB.Model(&model.Product{}).Where("id = ?", productID).First(&res).Error
}

func (r *productRepository) Update(productID int, req model.Product) error {
	return r.DB.Model(&model.Product{}).Where("id = ?", productID).Save(&req).Error
}
