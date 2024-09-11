package models

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	CampaignID uint     `json:"campaign_id" gorm:"not null"`
	Campaign   Campaign `json:"-"`
	Amount     float64  `json:"amount" gorm:"type:decimal(10,2);not null"`
	UserID     uint     `json:"user_id" gorm:"not null"`
	User       User     `json:"-"`
	Comment    string   `json:"comment" gorm:"type:text"`
}
