package dto

type UserLogoutRequest struct {
	Email string `json:"email" validate:"required,email"`
}
