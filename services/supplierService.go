package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utilities/request"
)

type SupplierService interface {
	CreateSupplier(createSupplierRequest request.CreateSupplierRequest, UserID uint) (model.Supplier, error)
}

type supplierService struct {
	repository repository.SupplierRepository
}

// CreateSupplier implements SupplierService.
func (s *supplierService) CreateSupplier(createSupplierRequest request.CreateSupplierRequest, UserID uint) (model.Supplier, error) {
	supplier := model.Supplier{
		CompanyName:    createSupplierRequest.CompanyName,
		CompanyAddress: createSupplierRequest.CompanyAddress,
		Contact:        createSupplierRequest.Contact,
		UserID:         UserID,
	}

	newSupplier, err := s.repository.CreateSupplier(supplier)
	return newSupplier, err
}

func NewSupplierService(repository repository.SupplierRepository) *supplierService {
	return &supplierService{repository}
}
