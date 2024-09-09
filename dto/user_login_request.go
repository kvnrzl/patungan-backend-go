package dto

import "bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLoginRequest) ToEntity() models.User {
	return models.User{
		Email:    u.Email,
		Password: u.Password,
	}
}
