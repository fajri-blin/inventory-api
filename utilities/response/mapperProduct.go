package response

import "inventory-api/model"

func ConvertToProductResponse(p model.Product) ProductResponse {
	return ProductResponse{
		ID:          int(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		SupplierID:  p.SupplierID,
	}
}
