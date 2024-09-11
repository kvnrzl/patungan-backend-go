package repositories

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DonationRepository struct {
}

func InitDonationRepository() DonationRepository {
	return DonationRepository{}
}

func (cr *DonationRepository) Create(tx *gorm.DB, input models.Donation) (models.Donation, error) {

	fmt.Println("input cuy :", input)
	fmt.Println("input user ID :", input.UserID)

	var donation models.Donation

	result := tx.Exec(`
        INSERT INTO donations (created_at, updated_at, deleted_at, campaign_id, amount, user_id, comment) 
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
		time.Now(),
		time.Now(),
		input.DeletedAt,
		input.CampaignID,
		input.Amount,
		input.UserID,
		input.Comment,
	)

	if err := result.Error; err != nil {
		return models.Donation{}, err
	}

	if err := result.Last(&donation).Error; err != nil {
		return models.Donation{}, err
	}

	return donation, nil
}

//func (cr *DonationRepository) Create(tx *gorm.DB, input models.Donation) (models.Donation, error) {
//
//	fmt.Println("input cuy :", input)
//	fmt.Println("input user ID :", input.UserID)
//
//	var donation models.Donation
//
//	result := tx.Create(&input)
//
//	if err := result.Error; err != nil {
//		return models.Donation{}, err
//	}
//
//	if err := result.Last(&donation).Error; err != nil {
//		return models.Donation{}, err
//	}
//
//	return donation, nil
//}
