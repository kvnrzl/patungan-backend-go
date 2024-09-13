package controllers

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"strconv"
)

type CampaignController struct {
	validate *validator.Validate
	service  services.CampaignService
}

func InitCampaignController(service services.CampaignService, validate *validator.Validate) CampaignController {
	return CampaignController{
		service:  service,
		validate: validate,
	}
}

func (cc *CampaignController) Create(c echo.Context) error {

	var createCampaignRequest dto.CreateCampaignRequest
	if err := c.Bind(&createCampaignRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := cc.validate.Struct(createCampaignRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	campaign := createCampaignRequest.ToEntity()
	campaign.FundraiserID = c.Get("userID").(uint)

	campaign, err := cc.service.Create(campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Campaign]{
		Status:  "success",
		Message: "campaign created",
		Data:    campaign,
	})
}

func (cc *CampaignController) GetAll(c echo.Context) error {
	campaigns, err := cc.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[[]models.Campaign]{
		Status:  "success",
		Message: "all campaigns",
		Data:    campaigns,
	})
}

func (cc *CampaignController) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	campaign, err := cc.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Campaign]{
		Status:  "success",
		Message: "campaign found",
		Data:    campaign,
	})
}

func (cc *CampaignController) GetByTitle(c echo.Context) error {
	title := c.Param("title")
	title, err := url.QueryUnescape(title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid title",
		})
	}

	campaign, err := cc.service.GetByTitle(title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Campaign]{
		Status:  "success",
		Message: "campaign found",
		Data:    campaign,
	})
}
