package request

type CreateSupplierRequest struct {
	CompanyName string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Contact string `json:"contact"`
}