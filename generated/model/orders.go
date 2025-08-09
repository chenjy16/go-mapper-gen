package model

import (
	"time"
)

// Orders 
type Orders struct {
	Id *int `db:"id" json:"id"`
	OrderNo string `db:"order_no" json:"order_no"`
	UserId int `db:"user_id" json:"user_id"`
	TotalAmount float64 `db:"total_amount" json:"total_amount"`
	Status *int `db:"status" json:"status"`
	PaymentMethod *int `db:"payment_method" json:"payment_method"`
	ShippingAddress *string `db:"shipping_address" json:"shipping_address"`
	Remark *string `db:"remark" json:"remark"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

// TableName 返回表名
func (Orders) TableName() string {
	return "orders"
}
