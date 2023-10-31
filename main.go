package main

import (
	"Project1/auth"
	"Project1/config"
	"Project1/model"
	"fmt"
)

func main() {
	var inputMenu = 0
	db, err := config.InitDB()

	if err != nil{
		fmt.Println("Cannot start program, database issue ", err.Error())
	}

	db.AutoMigrate(&model.User{})

	var auth = auth.AuthSystem{DB: db}
	for{
		fmt.Println("1. Login")
		fmt.Println("9. Exit")
		fmt.Scanln(&inputMenu)
		if inputMenu == 9 {
			break
		}else if inputMenu == 1{
			var inputDashboard int
			userLogin, permit := auth.Login()
			if permit {
				for{
					fmt.Println("========Dasboard========")
					fmt.Println("1. Input Barang")
					fmt.Println("2. Edit Barang")
					fmt.Println("3. Input Customer")
					fmt.Println("4. Input Transaksi")
					if userLogin.Role == 2 {
						fmt.Println("8. Menambahkan Pegawai baru")
					}
					fmt.Println("9. LogOut")
					fmt.Println("========================")
					fmt.Print("Masukkan pilihan: ")
					fmt.Scanln(&inputDashboard)
					if inputDashboard == 9 {
						permit = false
					}else if inputDashboard == 8 && userLogin.Role == 2{
						result, permit := auth.Register()
						if permit {
							fmt.Println(result)
						}
					}
				}
			}
			
		}
	}

}