package dto

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"time"
)

type CreateCampaignRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Target      float64   `json:"target"`
	EndDate     time.Time `json:"endDate" validate:"required"`
	CategoryID  uint      `json:"categoryID" validate:"required"`
}

// create json from the CreateCampaignRequest
/*
{
	"title": "Campaign 1",
	"description": "Description 1",
	"target": 1000000,
	"endDate": "2021-08-31T00:00:00Z",
	"categoryID": 1
}
*/

// create function to convert CreateCampaignRequest to Campaign
func (ccr *CreateCampaignRequest) ToEntity() models.Campaign {
	return models.Campaign{
		Title:       ccr.Title,
		Description: ccr.Description,
		Target:      ccr.Target,
		EndDate:     ccr.EndDate,
		CategoryID:  ccr.CategoryID,
	}
}
