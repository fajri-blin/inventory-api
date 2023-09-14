package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier model.Supplier) (model.Supplier, error)
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *supplierRepository {
	return &supplierRepository{db}
}

