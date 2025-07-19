package dto

type CreateProductRequest struct {
	Name  string `json:"name" binding:"required" validate:"required"`
	Stock int64  `json:"stock" binding:"required" validate:"required"`
}
