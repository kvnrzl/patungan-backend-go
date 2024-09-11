package controllers

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PaymentController struct {
	validate *validator.Validate
	service  services.PaymentService
}

func InitPaymentController(service services.PaymentService, validate *validator.Validate) PaymentController {
	return PaymentController{
		service:  service,
		validate: validate,
	}
}

func (dc *PaymentController) CreatePayment(c echo.Context) error {

	var createPaymentRequest dto.CreatePaymentRequest
	if err := c.Bind(&createPaymentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := dc.validate.Struct(createPaymentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	payment := createPaymentRequest.ToEntity()
	logrus.Println("payment: ", payment)

	payment, snapResp, err := dc.service.CreateTransaction(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	newData := struct {
		Payment      models.Payment `json:"payment"`
		SnapResponse *snap.Response `json:"snap_response"`
	}{
		Payment:      payment,
		SnapResponse: snapResp,
	}

	return c.JSON(http.StatusOK, models.BaseResponse[any]{
		Status:  "success",
		Message: "payment created",
		Data:    newData,
	})
}
