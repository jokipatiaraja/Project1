package transaction

import (
	"Project1/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)


type AuthSystem struct {
	DB *gorm.DB
}


func (as *AuthSystem) Add(userId uint){

	var inputMenu = 0

	isTransaction := true
	for isTransaction {
		fmt.Println("1. LIhat data customer")
		fmt.Println("2. Input transaksi")
		fmt.Println("3. Exit")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1{
			var listCustomer = new([]model.Customer)
			result := as.DB.Find(&listCustomer)
			if result.Error != nil {
				fmt.Println("Something wrong : ", result.Error)
				continue
			}
			fmt.Println("###### List Customer ######")
			for _, customer := range *listCustomer {
				fmt.Println(customer.ID, "-",customer.Name)
			}
			fmt.Println("#########################")
		}else if inputMenu == 2{
			var newTransactionDetail = new(model.TransactionDetails)
			var newTransaction = new(model.Transaction)
			var	idCustomer uint = 0
			var transactionDetails = []model.TransactionDetails{}
			newTransaction.UserID = userId
			var customer = new(model.Customer)
			var product = new(model.Product)
			
			
			fmt.Print("Masukkan ID Customer : ")
			fmt.Scanln(&idCustomer)
			fmt.Println(idCustomer)
			var err = as.DB.First(&customer,idCustomer).Error
			if err != nil{
				fmt.Println("Data Customer tidak di temukan, periksa kembali Id Customer")
				continue
			}

			isBuying := true
			totalPayment := 0
			totalQty := 0
			for isBuying{
				fmt.Print("Masukkan ID product yang ingin dibeli")
				fmt.Scanln(&newTransactionDetail.ProductID)

				fmt.Println(newTransactionDetail.ProductID)
				var err = as.DB.First(&product,newTransactionDetail.ProductID).Error
				if err != nil{
					fmt.Println("Data Barang tidak di temukan, periksa kembali Id Barang")
					continue
				}
				fmt.Print("Masukkan jumlah barang yang dibeli : ")
				fmt.Scanln(&newTransactionDetail.Qty)
				totalQty += int(newTransactionDetail.Qty)
				newTransactionDetail.Price = product.Price * newTransactionDetail.Qty
				totalPayment += int(newTransactionDetail.Price)
				newTransactionDetail.Created_at = time.Now()
				
				transactionDetails = append(transactionDetails, *newTransactionDetail)
				totalPayment += int(newTransactionDetail.Price)
				belilagi := 0
				fmt.Print("Ketik angka 1 untuk input lagi, dan angka lain untuk selesai input transaksi : ")
				fmt.Scanln(&belilagi)
				if belilagi != 1 {
					isBuying = false
				}
			}
			newTransaction.TotalPrice = uint(totalPayment)
			newTransaction.TotalQty = uint16(totalQty)
			newTransaction.CustomerId = idCustomer
			
			errTransaction := as.DB.Create(newTransaction).Error

			if err != nil{
				fmt.Println("Something wrong  transaction : ", errTransaction)
			}
			for i := range transactionDetails{
				transactionDetails[i].TransactionID = newTransaction.ID
				fmt.Println(transactionDetails[i])
				errTransactionDetail := as.DB.Create(transactionDetails[i]).Error
				if errTransactionDetail != nil {
					fmt.Println("Something wrong transaction detail : ", errTransactionDetail)
					break
				}
			}

		}else if inputMenu == 3{
			isTransaction = false
		}
	}
	// var newTransactionDetail = new(model.TransactionDetails)
	// var newTransaction = new(model.Transaction)



	// fmt.Print("Masukkan Nama Product		: ")
	// fmt.Scanln(&newProduct.Name)
	// fmt.Print("Masukkan Jumlah Stock		: ")
	// fmt.Scanln(&newProduct.Stock)
	// fmt.Print("Masukkan Harga Product		: ")
	// fmt.Scanln(&newProduct.Price)
	// newProduct.UserID = userId

	// err := as.DB.Create(newProduct).Error

	// if err != nil {
	// 	fmt.Println("input error:", err.Error())
	// 	return model.Product{}, false
	// }

	// return *newProduct, true
}
