package models

import (
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	gorm.Model
	FundraiserID uint      `json:"fundraiser_id" gorm:"not null"`
	User         User      `json:"-" gorm:"foreignKey:FundraiserID;references:ID"`
	Title        string    `json:"title" gorm:"type:varchar(100); uniqueIndex; not null"`
	Description  string    `json:"description" gorm:"type:text; not null"`
	Collected    float64   `json:"collected" gorm:"type:decimal(10,2); default:0"`
	Target       float64   `json:"target" gorm:"type:decimal(10,2); not null"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Status       string    `json:"status" gorm:"type:varchar(20); default:'active'"`
	CategoryID   uint      `json:"category_id" gorm:"not null"`
	Category     Category  `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}
