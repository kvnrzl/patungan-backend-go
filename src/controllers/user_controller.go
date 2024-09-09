package controllers

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	service  services.UserService
	validate *validator.Validate
}

func InitUserController(service services.UserService, validate *validator.Validate) UserController {
	return UserController{
		validate: validate,
		service:  service,
	}
}

func (uc *UserController) LoginHandler(c echo.Context) error {

	// get user from request
	var userLoginRequest dto.UserLoginRequest
	if err := c.Bind(&userLoginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	// validate request
	err := uc.validate.Struct(userLoginRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	user := userLoginRequest.ToEntity()

	// call service
	token, err := uc.service.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}

func (uc *UserController) RegisterHandler(c echo.Context) error {

	// get user from request
	var userRegisterRequest dto.UserRegisterRequest
	if err := c.Bind(&userRegisterRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	// validate request
	err := uc.validate.Struct(userRegisterRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	user := userRegisterRequest.ToEntity()

	// call service
	newUser, err := uc.service.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.User]{
		Status:  "success",
		Message: "register success",
		Data:    newUser,
	})

}
