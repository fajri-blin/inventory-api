package request

type ProductRequest struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price int `json:"price" binding:"required,number"`
	Quantity int `json:"quantity" binding:"required,number"`
	SupplierID uint `json:"suplier_id" binding:"required"`
}