package services

import (
	"errors"
	"inventory-api/model"
	"inventory-api/repository"
	"inventory-api/utilities/request"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(signupRequest request.SignUpRequest) (model.User, error)
	Login(loginRequest request.LoginRequest) (string, error)
	DeleteUser(ID int) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUser(signupRequest request.SignUpRequest) (model.User, error) {
	//Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	//Save User
	user := model.User{
		Email:      signupRequest.Email,
		Password:   string(hash),
		IsSupplier: false,
	}

	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}

func (s *userService) DeleteUser(ID int) (model.User, error) {
	user, err := s.repository.FindByID(ID)
	_, err = s.repository.DeleteUser(user)

	return user, err
}

func (s *userService) Login(loginRequest request.LoginRequest) (string, error) {
	//get user
	user, err := s.repository.FindByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	} else if user.ID == 0 {
		return "", errors.New("invalid email or password")
	}

	//compared password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", err
	}

	//sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        user.ID,
		"isSupplier": user.IsSupplier,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
