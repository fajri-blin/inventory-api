package request

type CreateSupplierRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	CompanyAddress string `json:"company_address" binding:"required"`
	Contact string `json:"contact" binding:"required"`
}