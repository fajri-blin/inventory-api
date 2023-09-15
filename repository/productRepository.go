package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(ID int) (model.Product, error)
	FindBySupplierID(supplierID int) ([]model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(product model.Product) (model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(ID int) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, ID).Error
	return product, err
}

// Find Product by Supplier ID
func (r *productRepository) FindBySupplierID(supplierID int) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("supplier_id = ?", supplierID).Find(&products).Error
	return products, err
}

// Create
func (r *productRepository) Create(product model.Product) (model.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

// Update
func (r *productRepository) Update(product model.Product) (model.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

// Delete
func (r *productRepository) Delete(product model.Product) (model.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
