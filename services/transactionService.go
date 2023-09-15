package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utilities/request"
)

type TransactionService interface {
	FindAll() ([]model.Transaction, error)
	FindByUserID(UserID uint) ([]model.Transaction, error)
	FindByID(ID int) (model.Transaction, error)
	FindBySupplierID(SupplierID int) ([]model.Transaction, error)
	Create(transactionRequest request.CreateTransaction, UserID uint) (model.Transaction, error)
	Update(ID int, transactionRequest request.UpdateTransaction) (model.Transaction, error)
	Delete(ID int) (model.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

// Delete implements TransactionService.
func (*transactionService) Delete(ID int) (model.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) FindAll() ([]model.Transaction, error) {
	trx, err := s.transactionRepository.FindAll()
	return trx, err
}

func (s *transactionService) FindByUserID(UserID uint) ([]model.Transaction, error) {
	trx, err := s.transactionRepository.FindByUserID(UserID)
	return trx, err
}

func (s *transactionService) FindByID(ID int) (model.Transaction, error) {
	trx, err := s.transactionRepository.FindByID(ID)
	return trx, err
}

//Find transactions by SupplierID
func (s *transactionService) FindBySupplierID(SupplierID int) ([]model.Transaction, error) {
	transactions, err := s.transactionRepository.FindBySupplierID(SupplierID)
	return transactions, err
}

func (s *transactionService) Create(transaction request.CreateTransaction, UserID uint) (model.Transaction, error) {
	trx := model.Transaction{
		Type:       transaction.Type,
		UserID:     UserID,
		ProductID:  transaction.ProductID,
		SupplierID: transaction.SupplierID,
		Quantity:   transaction.Quantity,
	}

	newTrx, err := s.transactionRepository.Create(trx)
	return newTrx, err
}

func (s *transactionService) Update(ID int, transactionRequest request.UpdateTransaction) (model.Transaction, error) {
	trx, _ := s.transactionRepository.FindByID(ID)

	if transactionRequest.Type != "" || transactionRequest.ProductID != 0 || transactionRequest.SupplierID != 0 || transactionRequest.Quantity != 0 {
		trx.Type = transactionRequest.Type
		trx.ProductID = transactionRequest.ProductID
		trx.SupplierID = transactionRequest.SupplierID
		trx.Quantity = transactionRequest.Quantity
	}

	newTrx, err := s.transactionRepository.Update(trx)
	return newTrx, err
}
