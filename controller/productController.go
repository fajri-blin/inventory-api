package controller

import "inventory-api/services"

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *productController {
	return &productController{productService}
}

// Create

