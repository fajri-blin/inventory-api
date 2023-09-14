package response

import "inventory-api/model"

type TransactionResponse struct {
	ID         uint   `json:"id"`
	Type       string `json:"type"`
	ProductID  uint   `json:"product_id"`
	SupplierID uint   `json:"supplier_id"`
	Quantity   int    `json:"quantity"`
}

func ConvertToTransactionResponse(transaction model.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:         transaction.ID,
		Type:       transaction.Type,
		ProductID:  transaction.ProductID,
		SupplierID: transaction.SupplierID,
		Quantity:   transaction.Quantity,
	}
}
