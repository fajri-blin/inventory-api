package repository

import{
	"inventory-api/model"
}

type UserRepository interface {
	Create(user model.User) (model.User, error)
}

type userRepository struct {
	db
}