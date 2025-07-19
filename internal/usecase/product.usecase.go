package usecase

import (
	"errors"
	"stock-and-order-management/internal/domain"
	"stock-and-order-management/internal/dto"
	"stock-and-order-management/internal/model"
	sharedError "stock-and-order-management/shared/error"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (u *productUsecase) Create(req dto.CreateProductRequest) error {
	product := &model.Product{}
	err := u.repo.CheckDuplicate(req.Name, product)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if err == nil {
		return errors.New(sharedError.ERR_CONFLICT)
	}

	product = &model.Product{}
	if err := copier.Copy(&product, &req); err != nil {
		return err
	}

	now := time.Now()
	product.CreatedAt = &now
	product.UpdatedAt = &now

	return u.repo.Create(*product)

}
