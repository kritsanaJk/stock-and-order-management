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

type orderUsecase struct {
	repo        domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewOrderUsecase(repo domain.OrderRepository, productRepo domain.ProductRepository) domain.OrderUsecase {
	return &orderUsecase{
		repo:        repo,
		productRepo: productRepo,
	}
}

func (u *orderUsecase) Create(req dto.CreateOrderRequest) error {
	product := &model.Product{}
	err := u.productRepo.GetByID(int(req.ProductID), product)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(sharedError.ERR_BAD_REQUEST)
		}
		return err
	}

	if req.Quantity > product.Stock {
		return errors.New("QUANTITY_NOT_ENOUGH")
	}

	order := &model.Order{}
	if err := copier.Copy(&order, &req); err != nil {
		return err
	}

	now := time.Now()
	order.CreatedAt = &now
	order.IdempotencyKey = now.GoString()
	if err := u.repo.Create(*order); err != nil {
		return err
	}

	product.UpdatedAt = &now
	product.Stock = product.Stock - req.Quantity
	return u.productRepo.Update(int(req.ProductID), *product)

}
