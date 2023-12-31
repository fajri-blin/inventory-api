package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id int) (model.User, error) {
	var user model.User

	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *userRepository) DeleteUser(user model.User) (model.User, error) {
	err := r.db.Delete(user).Error

	return user, err
}
