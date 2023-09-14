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
	"github.com/golang-jwt/jwt/v5"
)

type supplierController struct {
	supplierService services.SupplierService
}

func NewSupplierController(supplierService services.SupplierService) *supplierController {
	return &supplierController{supplierService}
}

func (h *supplierController) GetAllSupplier(c *gin.Context){
	suppliers, err := h.supplierService.FindAllSupplier()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var supplierResponses []response.SupplierResponse

	for _, supplier := range suppliers{
		supplierResponse := response.ConvertToSupplierResponseHandler(supplier)
		supplierResponses = append(supplierResponses, supplierResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": supplierResponses,
	})
}

func (h *supplierController) GetSupplierByID(c *gin.Context){
	ID, _ := strconv.Atoi(c.Param("id"))
	supplier, err := h.supplierService.FindSupplierByID(ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}

	supplierResponse := response.ConvertToSupplierResponseHandler(supplier)
	c.JSON(http.StatusOK, gin.H{
		"data": supplierResponse,
	})
}

func (h *supplierController) CreateCompanyController(c *gin.Context){
	var supplierRequest request.CreateSupplierRequest

	err := c.ShouldBindJSON(&supplierRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return

		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)

	fmt.Println("User ID : ", userID)

	supplier, err := h.supplierService.CreateSupplier(supplierRequest, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": supplier,
	})
}

func (h *supplierController) UpdateSupplier(c *gin.Context){
	var supplierRequest request.UpdateSupplierRequest
	err := c.ShouldBindJSON(&supplierRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return

		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
	}

	ID, _ := strconv.Atoi(c.Param("id"))
	s, err := h.supplierService.UpdateSupplier(ID, supplierRequest)
	supplierResponse := response.ConvertToSupplierResponseHandler(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": supplierResponse,
	})
}

func (h *supplierController) DeleteUser(c *gin.Context){
	ID, _ :=strconv.Atoi(c.Param("id"))
	a, err := h.supplierService.DeleteSupplier(ID)
	supplierResponse := response.ConvertToSupplierResponseHandler(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": supplierResponse,
	})
}

