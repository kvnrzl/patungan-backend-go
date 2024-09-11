package dto

import "bitbucket.org/bri_bootcamp/patungan-backend-go/models"

type CreatePaymentRequest struct {
	DonationID    uint    `json:"donation_id" validate:"required"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount" validate:"required"`
	Currency      string  `json:"currency"`
}

// create json from the CreatePaymentRequest
/*
{
	"donation_id": 1,
	"payment_method": "OVO",
	"amount": 1000000,
	"currency": "IDR"
}
*/

func (r *CreatePaymentRequest) ToEntity() models.Payment {
	return models.Payment{
		DonationID:    r.DonationID,
		PaymentMethod: r.PaymentMethod,
		Amount:        r.Amount,
		Currency:      r.Currency,
	}
}
