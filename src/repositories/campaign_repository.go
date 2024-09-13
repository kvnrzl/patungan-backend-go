package repositories

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"fmt"
	"gorm.io/gorm"
)

type CampaignRepository struct {
}

func InitCampaignRepository() CampaignRepository {
	return CampaignRepository{}
}

func (cr *CampaignRepository) GetAll(tx *gorm.DB) ([]models.Campaign, error) {
	var campaigns []models.Campaign

	if err := tx.Preload("Category").Find(&campaigns).Error; err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (cr *CampaignRepository) GetByID(tx *gorm.DB, id uint) (models.Campaign, error) {
	var campaign models.Campaign

	if err := tx.Preload("Category").First(&campaign, "id = ?", id).Error; err != nil {
		return models.Campaign{}, err
	}

	return campaign, nil
}

func (cr *CampaignRepository) GetByTitle(tx *gorm.DB, title string) (models.Campaign, error) {
	var campaign models.Campaign

	fmt.Println("database DB :", tx)
	if err := tx.Preload("Category").First(&campaign, "title = ?", title).Error; err != nil {
		return models.Campaign{}, err
	}

	return campaign, nil
}

func (cr *CampaignRepository) Create(tx *gorm.DB, campaignInput models.Campaign) (models.Campaign, error) {

	var campaign models.Campaign

	result := tx.Create(&campaignInput)

	if err := result.Error; err != nil {
		return models.Campaign{}, err
	}

	if err := result.Last(&campaign).Error; err != nil {
		return models.Campaign{}, err
	}

	return campaign, nil
}

//func (cr *CampaignRepository) Update(tx *gorm.DB, campaignInput models.Campaign, id uint) (models.Campaign, error) {
//	campaign, err := cr.GetByID(tx, id)
//
//	if err != nil {
//		return models.Campaign{}, err
//	}
//
//	campaign.Name = campaignInput.Name
//
//	if err := tx.Save(&campaign).Error; err != nil {
//		return models.Campaign{}, err
//	}
//
//	return campaign, nil
//}

func (cr *CampaignRepository) UpdateCollected(tx *gorm.DB, input models.Campaign) (models.Campaign, error) {
	var campaign models.Campaign

	campaign = input
	if err := tx.Save(&campaign).Error; err != nil {
		return models.Campaign{}, err
	}

	return campaign, nil
}

func (cr *CampaignRepository) Delete(tx *gorm.DB, id uint) error {
	campaign, err := cr.GetByID(tx, id)

	if err != nil {
		return err
	}

	if err := tx.Delete(&campaign).Error; err != nil {
		return err
	}

	return nil
}
