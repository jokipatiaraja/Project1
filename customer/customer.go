package customer

import (
	"Project1/model"
	"fmt"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) add() (model.Customer, bool){
	var newCustomer = new(model.Customer)

	fmt.Print("Masukkan Nama		:")
	fmt.Scanln(&newCustomer.Name)
	fmt.Print("Masukkan User Name	:")
	fmt.Scanln(&newCustomer.UserName)
	fmt.Print("Masukkan Password	:")
	fmt.Scanln(&newCustomer.Password)

	err := as.DB.Create(newCustomer).Error
	if err != nil {
		fmt.Println("Input error:", err.Error())
		return model.Customer{}, false
	}

	return *newCustomer, true
}