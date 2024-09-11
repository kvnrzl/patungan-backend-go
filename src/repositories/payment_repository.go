package repositories

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
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
