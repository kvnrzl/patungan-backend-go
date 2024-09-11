package repositories

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"fmt"
	"gorm.io/gorm"
)

type DonationRepository struct {
}

func InitDonationRepository() DonationRepository {
	return DonationRepository{}
}

func (cr *DonationRepository) Create(tx *gorm.DB, input models.Donation) (models.Donation, error) {

	var donation models.Donation

	//result := tx.Exec(`
	//    INSERT INTO donations (created_at, updated_at, deleted_at, campaign_id, amount, user_id, comment)
	//    VALUES (?, ?, ?, ?, ?, ?, ?)`,
	//	time.Now(),
	//	time.Now(),
	//	input.DeletedAt,
	//	input.DonationID,
	//	input.Amount,
	//	input.UserID,
	//	input.Comment,
	//)

	result := tx.Create(&input)

	if err := result.Error; err != nil {
		return models.Donation{}, err
	}

	if err := result.Last(&donation).Error; err != nil {
		return models.Donation{}, err
	}

	return donation, nil
}

// create function get by id
func (cr *DonationRepository) GetByID(tx *gorm.DB, id uint) (models.Donation, error) {
	var donation models.Donation

	fmt.Println("database DB :", tx)
	if err := tx.First(&donation, "id = ?", id).Error; err != nil {
		return models.Donation{}, err
	}

	return donation, nil
}
