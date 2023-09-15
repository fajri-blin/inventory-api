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

type transactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(service services.TransactionService) *transactionController {
	return &transactionController{service}
}

func (trx *transactionController) Create(c *gin.Context) {
	var transactionRequest request.CreateTransaction
	err := c.ShouldBindJSON(&transactionRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMsgs := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMsg := fmt.Sprintf("Error on field %s. %s", e.Field(), e.ActualTag())
				errorMsgs = append(errorMsgs, errorMsg)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMsgs,
			})
			return

		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)

	transaction, err := trx.transactionService.Create(transactionRequest, uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

// GetAll
func (trx *transactionController) GetAll(c *gin.Context) {
	transaction, err := trx.transactionService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if len(transaction) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":"There's no Transactions in Database",
		})
		return
	}

	var transactionsResponse []response.TransactionResponse
	for _, transaction := range transaction {
		transactionResponse := response.ConvertToTransactionResponse(transaction)
		transactionsResponse = append(transactionsResponse, transactionResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsResponse,
	})
}

//Find Transactions By SupplierID
func (trx *transactionController) FindBySupplierID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	transaction, err := trx.transactionService.FindBySupplierID(int(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if len(transaction) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":"There's no Transactions in Database",
		})
		return
	}
	
	var transactionsResponse []response.TransactionResponse
	for _, transaction := range transaction {
		transactionResponse := response.ConvertToTransactionResponse(transaction)
		transactionsResponse = append(transactionsResponse, transactionResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsResponse,
	})
}

// Get ByID
func (transaction *transactionController) GetByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	data, err := transaction.transactionService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if data.ID == 0 {
		result := fmt.Sprintf("Transaction with ID %d not found", ID)
		c.JSON(http.StatusNotFound, gin.H{
			"error" : result,
		})
	}

	transactionResponse := response.ConvertToTransactionResponse(data)

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResponse,
	})
}

// Update
func (trx *transactionController) Update(c *gin.Context) {
	var transactionRequest request.UpdateTransaction

	err := c.ShouldBindJSON(&transactionRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"erros": err.Error(),
			})
			return
		}
	}

	ID, _ := strconv.Atoi(c.Param("id"))
	data, err := trx.transactionService.Update(ID, transactionRequest)
	if data.ID == 0 {
		resultErr := fmt.Sprintf("Transaction with ID %d not found", ID)
			c.JSON(http.StatusNotFound, gin.H{
				"error": resultErr,
			})
		return
	}

	transactionResponse := response.ConvertToTransactionResponse(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResponse,
	})
}

// Delete
func (trx *transactionController) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	data, err := trx.transactionService.Delete(ID)

	if data.ID == 0 {
		result := fmt.Sprintf("Transaction with ID %d not found", ID)
		c.JSON(http.StatusNotFound, gin.H{
			"error": result,
		})
	}


	transactionResponse := response.ConvertToTransactionResponse(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResponse,
	})
}
