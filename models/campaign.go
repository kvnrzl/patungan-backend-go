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
	Target       float64   `json:"target" gorm:"type:decimal(10,2); not null"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Status       string    `json:"status" gorm:"type:varchar(20); default:'active'"`
	CategoryID   uint      `json:"category_id" gorm:"not null"`
	Category     Category  `json:"-"`
}
