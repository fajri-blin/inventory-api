package request

type CreateTransaction struct {
	Type       string `json:"type" binding:"required"`
	ProductID  uint   `json:"product_id" binding:"required,number"`
	SupplierID uint   `json:"supplier_id" binding:"required,number"`
	Quantity   int    `json:"quantity" binding:"required,number"`
}

type UpdateTransaction struct {
	Type       string `json:"type" binding:"required"`
	ProductID  uint   `json:"product_id" binding:"required,number"`
	SupplierID uint   `json:"supplier_id" binding:"required,number"`
	Quantity   int    `json:"quantity" binding:"required,number"`
}
