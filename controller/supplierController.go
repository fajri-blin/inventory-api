package controller

import (
	"encoding/json"
	"fmt"
	"inventory-api/services"
	"inventory-api/utilities/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type supplierController struct {
	supplierService services.SupplierService
}

func NewSupplierController(supplierService services.SupplierService) *supplierController {
	return &supplierController{supplierService: supplierService}
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
	userID, _ := claims["id"].(float64)

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