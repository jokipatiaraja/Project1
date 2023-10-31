package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TotalPrice uint
	Status uint8
	TotalQty uint16
	UserID uint
	CustomerId uint
	Details []Product `gorm:"many2many:transaction_details;"`
}