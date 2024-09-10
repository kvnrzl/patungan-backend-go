package services

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/repositories"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type CampaignService struct {
	db         *gorm.DB
	repository repositories.CampaignRepository
}

func InitCampaignService(db *gorm.DB, repository repositories.CampaignRepository) CampaignService {
	return CampaignService{
		db:         db,
		repository: repository,
	}
}

func (cs *CampaignService) Create(campaign models.Campaign) (models.Campaign, error) {

	// set start campaign is default time now
	campaign.StartDate = time.Now()

	if campaign.EndDate.Before(campaign.StartDate) {
		logrus.Println("end date must be after start date")
		return models.Campaign{}, errors.New("end date must be after start date")
	}

	if campaign.Target < 10000 {
		logrus.Println("target must be greater than 10000")
		return models.Campaign{}, errors.New("target must be greater than 10000")
	}

	return cs.repository.Create(cs.db, campaign)
}

func (cs *CampaignService) GetAll() ([]models.Campaign, error) {
	return cs.repository.GetAll(cs.db)
}

func (cs *CampaignService) GetByID(id uint) (models.Campaign, error) {
	return cs.repository.GetByID(cs.db, id)
}

func (cs *CampaignService) GetByTitle(title string) (models.Campaign, error) {
	return cs.repository.GetByTitle(cs.db, title)
}
