package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utilities/request"
)

type SupplierService interface {
	Create(createSupplierRequest request.CreateSupplierRequest, UserID uint) (model.Supplier, error)
}

type supplierService struct {
	repository repository.SupplierRepository
}

func NewSupplierRepository(repository repository.SupplierRepository) *supplierService {
	return &supplierService{repository}
}

func (s *supplierService) Create(createSupplierRequest request.CreateSupplierRequest, UserID uint) (model.Supplier, error) {
	supplier := model.Supplier{
		CompanyName:    createSupplierRequest.CompanyName,
		CompanyAddress: createSupplierRequest.CompanyAddress,
		Contact:        createSupplierRequest.Contact,
		UserID:         UserID,
	}

	newSupplier, err := s.repository.Create(supplier)
	return newSupplier, err
}
