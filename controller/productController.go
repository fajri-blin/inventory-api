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

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *productController {
	return &productController{productService}
}

// GetAll
func (p *productController) GetAll(c *gin.Context) {
	products, err := p.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	//check if there's no products in database
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":"There's no products in database",
		})
		return
	}

	var productsResponse []response.ProductResponse
	for _,p := range products{
		productResponse := response.ConvertToProductResponse(p)
		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productsResponse,
	})
}

// Get ByID
func (p *productController) GetByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	pc,err := p.productService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if pc.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":"Product not found",
		})
		return
	}

	productResponse := response.ConvertToProductResponse(pc)

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

// Create
func (p *productController) Create(c *gin.Context) {
	var  productRequest request.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
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
				"errors": err.Error(),
			})
			return
		}
	}
	product, err := p.productService.Create(productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// Update
func (p *productController) Update(c *gin.Context) {
	var  productRequest request.ProductRequest

	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _,e := range err.(validator.ValidationErrors) {
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

	ID, _:= strconv.Atoi(c.Param("id"))
	pc, err := p.productService.Update(ID, productRequest)

	// check if there's no product on certain ID founded in the database
	if pc.ID == 0 {
		result := fmt.Sprintf("Product with ID: %d not found", ID)
		c.JSON(http.StatusNotFound, gin.H{
			"Error": result,
		})
		return
	}

	productResponse := response.ConvertToProductResponse(pc)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

// Delete
func (p *productController) Delete(c *gin.Context) {
	ID, _:= strconv.Atoi(c.Param("id"))
	pc, err := p.productService.Delete(ID)

	if pc.ID == 0 {
		result := fmt.Sprintf("Product with ID: %d not found", ID)
		c.JSON(http.StatusNotFound, gin.H{
			"error": result,
		})
		return
	}

	productResponse := response.ConvertToProductResponse(pc)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

