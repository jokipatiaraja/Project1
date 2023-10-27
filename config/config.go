package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func initDB() (*gorm.DB, error){
	dsn := "root:@tcp(localhost:3306)/database_toko?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("something happend", err)

	}
	return db, err
	
}