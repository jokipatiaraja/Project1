package model

import (
	"time"
)

type TransactionDetails struct {
	TransactionID 	uint `gorm:"primaryKey"`
	ProductID		uint `gorm:"primaryKey"`
	Price			uint
	Qty				uint
	Created_at 		time.Time
}