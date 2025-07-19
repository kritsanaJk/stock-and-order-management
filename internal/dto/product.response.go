package dto

type ProductResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" `
	Stock int64  `json:"stock"`
}
