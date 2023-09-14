package response

import "inventory-api/model"

type SupplierResponse struct {
	ID uint `json:"id"`
	CompanyName string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Contact string `json:"contact"`
}


func ConvertToSupplierResponseHandler(supplier model.Supplier) SupplierResponse {
	return SupplierResponse{
		ID: supplier.ID,
		CompanyName: supplier.CompanyName,
		CompanyAddress: supplier.CompanyAddress,
		Contact: supplier.Contact,
	}
}