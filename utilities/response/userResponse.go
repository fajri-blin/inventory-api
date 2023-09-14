package response

import "inventory-api/model"

type UserResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
}


func ConvertToUserResponseHandler(user model.User) UserResponse {
	return UserResponse{
		ID : user.ID,
		Email : user.Email, 
	}
}