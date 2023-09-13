package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"unique"`
	IsSupplier bool
	Password   string

	//Relationship
	Supplier     Supplier `constraint:OnUpdate:CASCADE,OnDelete:SET NULL`
	Transactions []Transaction
}
