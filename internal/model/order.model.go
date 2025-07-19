package model

import "time"

type Order struct {
	ID uint `gorm:"primary key;auto_increment" json:"id"`

	ProductID uint    `json:"productId"`
	Product   Product `gorm:"forienKey:ProductID; reference:ID" json:"product"`

	UserID         string     `gorm:"type:text; NOT NULL" json:"userId"`
	Quantity       int64      `gorm:"NOT NULL" json:"quantity"`
	Status         string     `gorm:"type:text; NOT NULL" json:"status"`
	IdempotencyKey string     `gorm:"type:text" json:"idempotencyKey"`
	CreatedAt      *time.Time `gorm:"default:now()" json:"createdAt"`
}

func (m *Order) TableName() string {
	return "orders"
}
