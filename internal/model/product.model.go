package model

import "time"

type Product struct {
	ID        uint       `gorm:"primary key;auto_increment" json:"id"`
	Name      string     `gorm:"type:text" json:"name" `
	Stock     int64      `gorm:"default:0" json:"stock"`
	CreatedAt *time.Time `gorm:"default:now()" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"default:now()" json:"updatedAt"`
}

func (m *Product) TableName() string {
	return "products"
}
