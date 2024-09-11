package controllers

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DonationController struct {
	validate *validator.Validate
	service  services.DonationService
}

func InitDonationController(service services.DonationService, validate *validator.Validate) DonationController {
	return DonationController{
		service:  service,
		validate: validate,
	}
}

func (dc *DonationController) CreateDonationRegisteredUser(c echo.Context) error {

	campaignIDStr := c.Param("campaign_id")
	campaignID, err := strconv.Atoi(campaignIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid campaign id",
		})
	}

	var createDonationRequest dto.CreateDonationRequest
	if err := c.Bind(&createDonationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := dc.validate.Struct(createDonationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	donation := createDonationRequest.ToEntity()
	logrus.Println("donation: ", donation)
	donation.UserID = c.Get("userID").(uint)
	donation.CampaignID = uint(campaignID)

	donation, err = dc.service.CreateDonationRegisteredUser(donation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Donation]{
		Status:  "success",
		Message: "donation created",
		Data:    donation,
	})
}

func (dc *DonationController) CreateDonationGuestUser(c echo.Context) error {

	campaignIDStr := c.Param("campaign_id")
	campaignID, err := strconv.Atoi(campaignIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid campaign id",
		})
	}

	var createDonationGuestRequest dto.CreateDonationGuestRequest
	if err := c.Bind(&createDonationGuestRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := dc.validate.Struct(createDonationGuestRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	donation := createDonationGuestRequest.ToEntity()
	donation.CampaignID = uint(campaignID)

	donation, err = dc.service.CreateDonationGuestUser(donation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Donation]{
		Status:  "success",
		Message: "donation created",
		Data:    donation,
	})

}
