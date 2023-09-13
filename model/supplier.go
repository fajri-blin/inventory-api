package model

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name string
	Address string

	//Relationship
	Product []Product
}