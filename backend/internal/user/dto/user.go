package dto

import "backend/internal/user/model"

type UserResponseDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func ToUserResponseDTO(user *model.User) UserResponseDTO {
	return UserResponseDTO{
		ID:       user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username: user.Username,
		Email:    user.Email,
	}
}
