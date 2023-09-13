package model

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	CompanyName    string
	CompanyAddress string
	Contact        string

	//Relationship
	UserID uint
	Products []Product
}
