package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utilities/request"
)

type ProductService interface {
	FindAll() ([]model.Product, error)
	FindByID(ID int) (model.Product, error)
	FindBySupplierID(supplierID int) ([]model.Product, error)
	Create(producRequest request.ProductRequest) (model.Product, error)
	Update(ID int, productRequest request.ProductRequest) (model.Product, error)
	Delete(ID int) (model.Product, error)
}

type service struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]model.Product, error) {
	products, err := s.repository.FindAll()
	return products, err
}

func (s *service) FindByID(ID int) (model.Product, error) {
	product, err := s.repository.FindByID(ID)
	return product, err
}

//Find Product By Supplier ID
func (s *service) FindBySupplierID(supplierID int) ([]model.Product, error) {
	products, err := s.repository.FindBySupplierID(supplierID)
	return products, err
}

// Create
func (s *service) Create(productRequest request.ProductRequest) (model.Product, error) {
	product := model.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Quantity:    productRequest.Quantity,
		SupplierID: productRequest.SupplierID,
	}

	newProduct, err := s.repository.Create(product)
	return newProduct, err
}

// Update
func (s *service) Update(ID int, productRequest request.ProductRequest) (model.Product, error) {
	product, err := s.repository.FindByID(ID)

	if productRequest.Name != "" {
		product.Name = productRequest.Name
	}
	if productRequest.Description != "" {
		product.Description = productRequest.Description
	}
	if productRequest.Price != 0 {
		product.Price = productRequest.Price
	}
	if productRequest.Quantity != 0 {
		product.Quantity = productRequest.Quantity
	}

	newProduct, err := s.repository.Update(product)
	return newProduct, err
}

// Delete
func (s *service) Delete(ID int) (model.Product, error) {
	product, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(product)

	return product, err
}
