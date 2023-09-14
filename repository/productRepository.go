package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Product, error)
	FindByID(ID int) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(product model.Product) (model.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindByID(ID int) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, ID).Error
	return product, err
}

// Create
func (r *repository) Create(product model.Product) (model.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

// Update
func (r *repository) Update(product model.Product) (model.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

// Delete
func (r *repository) Delete(product model.Product) (model.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
