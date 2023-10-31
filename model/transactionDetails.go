package model

import (
	"time"
)

type TransactionDetails struct {
	TransactionID 	uint `gorm:"primaryKey"`
	ProductID		uint `gorm:"primaryKey"`
	qty				uint
	created_at 		time.Time
	updated_at 		time.Time
}