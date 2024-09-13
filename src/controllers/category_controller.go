package controllers

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CategoryController struct {
	validate *validator.Validate
	service  services.CategoryService
}

func InitCategoryController(service services.CategoryService, validate *validator.Validate) CategoryController {
	return CategoryController{
		service:  service,
		validate: validate,
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	categories, err := cc.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[[]models.Category]{
		Status:  "success",
		Message: "all categories",
		Data:    categories,
	})
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	category, err := cc.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Category]{
		Status:  "success",
		Message: "category found",
		Data:    category,
	})
}
