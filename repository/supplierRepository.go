package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	CreateSupplier(supplier model.Supplier) (model.Supplier, error)
	Update(supplier model.Supplier) (model.Supplier, error)
	DeleteSupplier(supplier model.Supplier) (model.Supplier, error)
	FindAll() ([]model.Supplier, error)
	FindSupplierByID(id int) (model.Supplier, error)
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *supplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) CreateSupplier(supplier model.Supplier) (model.Supplier, error) {
	err := r.db.Create(&supplier).Error
	return supplier, err
}

func (r *supplierRepository) Update(supplier model.Supplier) (model.Supplier, error) {
	err := r.db.Save(&supplier).Error
	return supplier, err
}

func (r *supplierRepository) FindSupplierByID(id int) (model.Supplier, error) {
	var supplier model.Supplier

	err := r.db.Find(&supplier, id).Error
	return supplier, err
}

func (r *supplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	
	err := r.db.Find(&suppliers).Error

	return suppliers, err
}

func (r *userRepository) DeleteSupplier(supplier model.Supplier) (model.Supplier, error) {
	err := r.db.Delete(supplier).Error

	return supplier, err
}