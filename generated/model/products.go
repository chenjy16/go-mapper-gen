package model

import (
	"time"
)

// Products 
type Products struct {
	Id *int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	Price float64 `db:"price" json:"price"`
	Stock *int `db:"stock" json:"stock"`
	CategoryId *int `db:"category_id" json:"category_id"`
	Status *int `db:"status" json:"status"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

// TableName 返回表名
func (Products) TableName() string {
	return "products"
}
