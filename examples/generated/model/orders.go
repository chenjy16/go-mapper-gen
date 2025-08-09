package model

import (
	"time"
)

// Orders 
type Orders struct {
	Id *int `db:"id" json:"id"`
	UserId int `db:"user_id" json:"user_id"`
	TotalAmount float64 `db:"total_amount" json:"total_amount"`
	Status *string `db:"status" json:"status"`
	OrderDate *time.Time `db:"order_date" json:"order_date"`
	ShippingAddress *string `db:"shipping_address" json:"shipping_address"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

// TableName 返回表名
func (Orders) TableName() string {
	return "orders"
}
