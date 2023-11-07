package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TotalPrice uint
	TotalQty uint16
	UserID uint
	CustomerId uint
	Details []Product `gorm:"many2many:transaction_details;"`
}