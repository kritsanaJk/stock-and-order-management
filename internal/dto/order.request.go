package dto

type CreateOrderRequest struct {
	ProductID uint  `json:"productId" binding:"required" validate:"required"`
	Quantity  int64 `json:"quantity" binding:"required" validate:"required"`
}
