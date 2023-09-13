package services

import (
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utillities/request"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(signupRequest request.SignUpRequest) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Create(signupRequest request.SignUpRequest) (model.User, error) {
	//Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	//Save User
	user := model.User{
		Email:    signupRequest.Email,
		Password: string(hash),
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}
