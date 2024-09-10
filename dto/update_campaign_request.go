package dto

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"time"
)

type UpdateCampaignRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Target      float64   `json:"target" validate:"required"`
	EndDate     time.Time `json:"endDate"`
	CategoryID  uint      `json:"categoryID" validate:"required"`
}

func (ucr *UpdateCampaignRequest) ToEntity() models.Campaign {
	return models.Campaign{
		Title:       ucr.Title,
		Description: ucr.Description,
		Target:      ucr.Target,
		EndDate:     ucr.EndDate,
		CategoryID:  ucr.CategoryID,
	}
}
