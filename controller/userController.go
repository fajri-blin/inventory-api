package controller

import (
	"encoding/json"
	"fmt"
	"inventory-api/services"
	"inventory-api/utilities/request"
	"inventory-api/utilities/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (h *userController) DeleteUser(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	a, err := h.userService.DeleteUser(ID)

	if a.ID == 0 {
		resultErr := fmt.Sprintf("Account with ID %v not found", ID)
		c.JSON(http.StatusNotFound, gin.H{
			"error": resultErr,
		})
	}

	userResponse := response.ConvertToUserResponseHandler(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userController) SignUp(c *gin.Context) {
	var signupRequest request.SignUpRequest

	err := c.ShouldBindJSON(&signupRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	user, err := h.userService.CreateUser(signupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": fmt.Sprintf("User %v created successfully", user.Email),
	})
}

func (h *userController) Login(c *gin.Context) {
	var loginRequest request.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	tokenString, err := h.userService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"token": tokenString,
		},
	})
}
