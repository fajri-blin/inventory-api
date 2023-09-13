package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Transaction string
	Quantity int
	TransactionDate time.Time

	//Cardinality
	ProductID uint
	UserID uint
	SupplierID uint
}