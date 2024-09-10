package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name" gorm:"type:varchar(100); not null"`
	Email      string `json:"email" gorm:"type:varchar(100); uniqueIndex; not null"`
	Password   string `json:"password" gorm:"type:varchar(255)"`
	Phone      string `json:"phone" gorm:"type:varchar(20)"`
	IsVerified bool   `json:"is_verified" gorm:"type:boolean; not null"`
}
