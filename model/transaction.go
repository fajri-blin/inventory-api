package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Type            string
	Quantity        int
	TransactionDate time.Time

	//Relationship
	UserID     uint
	ProductID  uint
	SupplierID uint
}
