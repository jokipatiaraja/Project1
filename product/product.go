package product

import (
	"Project1/model"
	"fmt"

	"gorm.io/gorm"
)


type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) Add(userId uint) (model.Product, bool){

	var newProduct = new(model.Product)

	fmt.Print("Masukkan Nama Product		: ")
	fmt.Scanln(&newProduct.Name)
	fmt.Print("Masukkan Jumlah Stock		: ")
	fmt.Scanln(&newProduct.Stock)
	fmt.Print("Masukkan Harga Product		: ")
	fmt.Scanln(&newProduct.Price)
	newProduct.UserID = userId

	err := as.DB.Create(newProduct).Error

	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Product{}, false
	}

	return *newProduct, true
}

func (ts *AuthSystem) Update() {
	
	var inputMenu = 0

	isEdit := true
	for isEdit {
		fmt.Println("1. LIhat list Barang")
		fmt.Println("2. Edit Informasi Barang")
		fmt.Println("3. Updated Stock Barang")
		fmt.Println("4. Exit")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var listProduct = new([]model.Product)
			result := ts.DB.Find(&listProduct)
			if result.Error != nil {
				fmt.Println("Something wrong : ", result.Error)
				continue
			}
			fmt.Println("###### List Barang ######")
			for _, product := range *listProduct {
				fmt.Println(product.ID, "-",product.Name)
			}
			fmt.Println("#########################")
		}else if inputMenu == 2{
			var product = new(model.Product)
			var	idProduct uint = 0
			
			fmt.Print("Masukkan ID barang yang ingin di ubah")
			fmt.Scanln(&idProduct)

			result := ts.DB.Raw("SELECT * FROM products WHERE id=?", idProduct)
			if result.RowsAffected == 0{
				fmt.Println("Data Barang tidak di temukan, periksa kembali Id Barang")
				continue
			}
			fmt.Print("Masukkan Nama product yang baru")
			fmt.Scanln(&product.Name)
			fmt.Print("Masukkan harga yang baru")
			fmt.Scanln(&product.Price)
			err := ts.DB.Model(&product).Where("id = ?", idProduct).Updates(model.Product{Name: product.Name, Price: product.Price}).Error
			if err != nil {
				fmt.Println("input error:", err.Error())
				continue
			}
			fmt.Println("Data Barang berhasil di ubah")

		}else if inputMenu == 3{
			var product = new(model.Product)
			var	idProduct uint = 0
			
			fmt.Print("Masukkan ID barang yang ingin di ubah")
			fmt.Scanln(&idProduct)

			result := ts.DB.Raw("SELECT * FROM products WHERE id=?", idProduct)
			if result.RowsAffected == 0{
				fmt.Println("Data Barang tidak di temukan, periksa kembali Id Barang")
				continue
			}

			fmt.Print("Masukkan Stock product yang baru")
			fmt.Scanln(&product.Stock)
			err := ts.DB.Model(&product).Where("id = ?", idProduct).Update("stock", product.Stock).Error
			if err != nil {
				fmt.Println("input error:", err.Error())
				continue
			}
			
		}else if inputMenu == 4{
			isEdit = false
		}
	}
}