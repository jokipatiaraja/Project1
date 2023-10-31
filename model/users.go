package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	UserName string
	Password string
	Role uint8
	Product []Product
	Customer []Customer
}