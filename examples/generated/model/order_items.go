package model

import (
	"time"
)

// OrderItems 
type OrderItems struct {
	Id *int `db:"id" json:"id"`
	OrderId int `db:"order_id" json:"order_id"`
	ProductId int `db:"product_id" json:"product_id"`
	Quantity int `db:"quantity" json:"quantity"`
	UnitPrice float64 `db:"unit_price" json:"unit_price"`
	TotalPrice float64 `db:"total_price" json:"total_price"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

// TableName 返回表名
func (OrderItems) TableName() string {
	return "order_items"
}
