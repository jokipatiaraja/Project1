package product

import (
	"Project1/model"
	"fmt"

	"gorm.io/gorm"
)


type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) add(userId int) (model.Product, bool){

	var newProduct = new(model.Product)

	fmt.Print("Masukkan Nama Product		:")
	fmt.Scanln(&newProduct.Name)
	fmt.Print("Masukkan Jumlah Stock		:")
	fmt.Scanln(&newProduct.Stock)
	fmt.Print("Masukkan Harga Product		:")
	fmt.Scanln(&newProduct.Price)
	newProduct.UserID = uint(userId)

	err := as.DB.Create(newProduct).Error

	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Product{}, false
	}

	return *newProduct, true
}