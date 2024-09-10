package dto

import "bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

func (u *UserRegisterRequest) ToEntity() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
	}
}
