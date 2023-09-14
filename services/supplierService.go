package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utilities/request"
)

type SupplierService interface {
	CreateSupplier(createSupplierRequest request.CreateSupplierRequest, UserID uint) (model.Supplier, error)
	UpdateSupplier(ID int, supplierRequest request.UpdateSupplierRequest) (model.Supplier, error)
	DeleteSupplier(ID int) (model.Supplier, error)
}

type supplierService struct {
	repository repository.SupplierRepository
}


func NewSupplierService(repository repository.SupplierRepository) *supplierService {
	return &supplierService{repository}
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

func (s *supplierService) UpdateSupplier(ID int, supplierRequest request.UpdateSupplierRequest) (model.Supplier, error) {
	supplier, err := s.repository.FindSupplierByID(ID)

	if supplier.CompanyName != "" {supplier.CompanyName = supplierRequest.CompanyName}
	if supplier.CompanyAddress != "" {supplier.CompanyAddress = supplierRequest.CompanyAddress}
	if supplier.Contact != "" {supplier.Contact = supplierRequest.Contact}

	newSupplier, err := s.repository.Update(supplier)

	return newSupplier, err
}

func (s *supplierService) DeleteSupplier(ID int) (model.Supplier, error) {
	supplier, err := s.repository.FindSupplierByID(ID)
	_, err = s.repository.DeleteSupplier(supplier)

	return supplier, err
}