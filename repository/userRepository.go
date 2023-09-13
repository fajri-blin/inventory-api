package repository

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err:= r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (model.User, error){
	var user model.User

	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}