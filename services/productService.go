package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utillities/request"
)

type ProductService interface {
	FindAll() ([]model.Product, error)
	FindByID()(model.Product, error)
	Create()(model.Product, error)
	Update()(model.Product, error)
	Delete()(model.Product, error)
}

type service struct{
	repository repository.Repository
}

func NewRepository( repository repository.Repository) *service{
	return &service{repository}
}

func (s *service) FindAll() ([]model.Product, error) {
	products, err := s.repository.FindAll()
	return products, err
}

func (s *service) FindByID(ID int)(model.Product, error) {
	product, err := s.repository.FindByID(ID)
	return product, err
}

// Create
func (s *service) Create(productRequest request.ProductRequest)(model.Product, error) {
	product := model.Product{
		Name: productRequest.Name,
		Description: productRequest.Description,
		Price: productRequest.Price,
		Quantity: productRequest.Quantity,
	}

	newProduct, err := s.repository.Create(product)
	return newProduct, err
}

// Update
func (s *service) Update(ID int, productRequest request.ProductRequest)(model.Product, error) {
	product, err := s.repository.FindByID(ID)

	if productRequest.Name != ""  {
		product.Name = productRequest.Name
	}
	if productRequest.Description != ""  {
		product.Description = productRequest.Description
	}
	if productRequest.Price != 0  {
		product.Price = productRequest.Price
	}
	if productRequest.Quantity != 0  {
		product.Quantity = productRequest.Quantity
	}

	newProduct, err := s.repository.Update(product)
	return newProduct, err
}

// Delete
func (s *service) Delete(ID int)(model.Product, error) {
	product, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(product)

	return product, err
}
>>>>>>> Stashed changes
