package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string
	Stock *uint16
	Price uint
	UserID uint
}