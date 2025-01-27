package repositories

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"fmt"
	"gorm.io/gorm"
)

type PaymentRepository struct {
}

func InitPaymentRepository() PaymentRepository {
	return PaymentRepository{}
}

func (cr *PaymentRepository) GetByID(tx *gorm.DB, id uint) (models.Payment, error) {
	var payment models.Payment

	if err := tx.First(&payment, "id = ?", id).Error; err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func (cr *PaymentRepository) Create(tx *gorm.DB, userInput models.Payment) (models.Payment, error) {

	var payment models.Payment

	result := tx.Create(&userInput)

	if err := result.Error; err != nil {
		return models.Payment{}, err
	}

	if err := result.Last(&payment).Error; err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func (cr *PaymentRepository) Update(tx *gorm.DB, input models.Payment) (models.Payment, error) {
	fmt.Println("inputID : ", input.ID)
	payment, err := cr.GetByID(tx, input.ID)
	if err != nil {
		return models.Payment{}, err
	}

	payment.PaymentStatus = input.PaymentStatus
	payment.PaymentMethod = input.PaymentMethod
	payment.Amount = input.Amount

	if err := tx.Save(&payment).Error; err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

//func (cr *PaymentRepository) Update(tx *gorm.DB, userInput models.Payment) (models.Payment, error) {
//
//	var payment models.Payment
//
//	result := tx.Save(&userInput)
//
//	if err := result.Error; err != nil {
//		return models.Payment{}, err
//	}
//
//	if err := result.Last(&payment).Error; err != nil {
//		return models.Payment{}, err
//	}
//
//	return payment, nil
//}
