package controller

import (
	"encoding/json"
	"fmt"
	"inventory-api/services"
	"inventory-api/utilities/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type trxController struct {
	trxService services.TransactionService
}

func NewTransactionController(service services.TransactionService) *trxController {
	return &trxController{service}
}

func (h *trxController) PostTrxController(c *gin.Context) {
	var trxRequest request.CreateTransaction
	err := c.ShouldBindJSON(&trxRequest)

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

	trx, err := h.trxService.Create(trxRequest, uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": trx,
	})
}
