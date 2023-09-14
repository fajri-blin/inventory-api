package request

type CreateTransaction struct {
	Type       string `binding:"required"`
	ProductID  uint    `binding:"required,number"`
	SupplierID uint    `binding:"required,number"`
	Quantity   int    `binding:"required,number"`
}

type UpdateTransaction struct {
	Type       string
	ProductID  uint
	SupplierID uint
	Quantity   int
}
