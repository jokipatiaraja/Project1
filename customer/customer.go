package customer

import (
	"Project1/model"
	"fmt"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) Add(userId uint) (model.Customer, bool){
	var newCustomer = new(model.Customer)

	fmt.Print("Masukkan Nama		:")
	fmt.Scanln(&newCustomer.Name)
	newCustomer.UserID = userId

	err := as.DB.Create(newCustomer).Error
	if err != nil {
		fmt.Println("Input error:", err.Error())
		return model.Customer{}, false
	}

	return *newCustomer, true
}