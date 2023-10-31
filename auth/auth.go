package auth

import (
	"Project1/model"
	"fmt"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}


func (as *AuthSystem) Login() (model.User, bool) {

	var currentUser = new(model.User)

	fmt.Print("Masukkan Username : ")
	fmt.Scanln(&currentUser.UserName)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&currentUser.Password)

	qry := as.DB.Where("user_name = ? AND password = ?", currentUser.UserName, currentUser.Password).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("Login process error ", err.Error())
		return model.User{}, false
	}
	return *currentUser, true
}

func (as *AuthSystem) Register() (model.User, bool) {

	var newUser = new(model.User)
	fmt.Print("Masukkan Nama		:")
	fmt.Scanln(&newUser.Name)
	fmt.Print("Masukkan User Name	:")
	fmt.Scanln(&newUser.UserName)
	fmt.Print("Masukkan Password	:")
	fmt.Scanln(&newUser.Password)
	newUser.Role = 1
	// err := as.DB.Table("pelanggan").Create(newUser).Error
	err := as.DB.Create(newUser).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.User{}, false
	}

	return *newUser, true

}