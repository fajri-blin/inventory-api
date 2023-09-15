package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll() ([]model.Transaction, error)
	FindByUserID(UserID uint) ([]model.Transaction, error)
	FindByID(ID int) (model.Transaction, error)
	FindBySupplierID(SupplierID int) ([]model.Transaction, error)
	Create(transaction model.Transaction) (model.Transaction, error)
	Update(transation model.Transaction) (model.Transaction, error)
	Delete(transaction model.Transaction) (model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) FindAll() ([]model.Transaction, error) {
	var transaction []model.Transaction
	err := r.db.Find(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) FindByUserID(UserID uint) ([]model.Transaction, error) {
	var transaction []model.Transaction
	err := r.db.Where("user_id = ?", UserID).Find(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) FindByID(ID int) (model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.First(&transaction, ID).Error
	return transaction, err
}

//Find Transactions By SupplierID
func (r *transactionRepository) FindBySupplierID(SupplierID int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Where("supplier_id = ?", SupplierID).Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Create(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) Update(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) Delete(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}
