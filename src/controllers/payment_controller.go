package controllers

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/services"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
	"io"
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

func (dc *PaymentController) PaymentCallback(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		logrus.Println("Unable to read body")
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "Unable to read body",
		})
	}
	defer c.Request().Body.Close()

	var callbackData dto.MidtransCallback
	if err := json.Unmarshal(body, &callbackData); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	// Handle callback
	fmt.Printf("Transaction Status: %s\n", callbackData.TransactionStatus)
	fmt.Printf("Order ID: %s\n", callbackData.OrderID)

	// convert to payment model
	payment := callbackData.ToEntity()

	// You can update your database here based on callback
	err = dc.service.UpdatePaymentStatus(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: "failed to update payment status",
		})
	}

	return c.JSON(http.StatusOK, "Callback received")
}
