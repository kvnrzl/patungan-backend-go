package models

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	//ID                     uint      `json:"id" gorm:"primaryKey"`
	DonationID             uint       `json:"donation_id" gorm:"not null"`
	Donation               Donation   `json:"-"`
	PaymentMethod          string     `json:"payment_method" gorm:"type:varchar(45)"`                   // bank_transfer, credit_card, e-wallet
	PaymentStatus          string     `json:"payment_status" gorm:"type:varchar(20);default:'pending'"` // pending, failed, success
	Amount                 float64    `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency               string     `json:"currency" gorm:"type:varchar(5);default:'IDR'"`
	TransactionID          string     `json:"transaction_id" gorm:"type:varchar(100)"`
	PaymentGatewayResponse string     `json:"payment_gateway_response,omitempty" gorm:"type:text"`
	PaidAt                 *time.Time `json:"paid_at"`
	PaymentChannel         string     `json:"payment_channel" gorm:"type:varchar(50)"` // bri, bca, gopay, ovo
	FailedReason           string     `json:"failed_reason" gorm:"type:text"`
	//CreatedAt              time.Time `json:"created_at"`
	//UpdatedAt              time.Time `json:"updated_at"`
}
