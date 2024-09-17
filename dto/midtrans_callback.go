package dto

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"gorm.io/gorm"
	"strconv"
)

type MidtransCallback struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	GrossAmount       string `json:"gross_amount"`
}

// to models.Payment
func (mc *MidtransCallback) ToEntity() models.Payment {
	paymentID, err := strconv.Atoi(mc.OrderID)
	if err != nil {
		return models.Payment{}
	}

	// convert string to float64
	amount, err := strconv.ParseFloat(mc.GrossAmount, 64)

	return models.Payment{
		Model: gorm.Model{
			ID: uint(paymentID),
		},
		PaymentStatus: mc.TransactionStatus,
		PaymentMethod: mc.PaymentType,
		Amount:        amount,
	}
}
