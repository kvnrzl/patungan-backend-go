package dto

import "bitbucket.org/bri_bootcamp/patungan-backend-go/models"

type CreateDonationRequest struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Amount  float64 `json:"amount" validate:"required"`
	Comment string  `json:"comment" validate:"required"`
}

func (r *CreateDonationRequest) ToEntity() models.Donation {
	return models.Donation{
		Amount:  r.Amount,
		Comment: r.Comment,
	}
}
