package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       int
	Quantity    int

	// Relationship
	SupplierID uint
	Transactions []Transaction
}
