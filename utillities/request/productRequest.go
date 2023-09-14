package request

type ProductRequest struct {
	Name string `binding:"required"`
	Description string `binding:"required"`
	Price int `json:"price" binding:"required,number"`
	Quantity int `json:"quantity" binding:"required,number"`
}