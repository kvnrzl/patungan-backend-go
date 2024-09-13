package dto

import "bitbucket.org/bri_bootcamp/patungan-backend-go/models"

type CreateDonationGuestRequest struct {
	Name    string  `json:"name" validate:"required"`
	Email   string  `json:"email" validate:"required,email"`
	Phone   string  `json:"phone" validate:""`
	Amount  float64 `json:"amount" validate:"required"`
	Comment string  `json:"comment" validate:"required"`
}

func (r *CreateDonationGuestRequest) ToEntity() models.Donation {
	return models.Donation{
		User: models.User{
			Name:  r.Name,
			Email: r.Email,
			Phone: r.Phone,
		},
		Amount:  r.Amount,
		Comment: r.Comment,
	}
}
