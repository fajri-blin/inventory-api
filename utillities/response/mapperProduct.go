package response

import "inventory-api/model"

func ConvertToProductResponse(p model.Product) ProductResponse{
	return ProductResponse{
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
		Quantity: p.Quantity,
	}
}