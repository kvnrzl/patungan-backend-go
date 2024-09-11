package services

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/repositories"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentService struct {
	db                 *gorm.DB
	paymentRepository  repositories.PaymentRepository
	donationRepository repositories.DonationRepository
}

func InitPaymentService(db *gorm.DB, paymentRepository repositories.PaymentRepository, donationRepository repositories.DonationRepository) PaymentService {
	return PaymentService{
		db:                 db,
		paymentRepository:  paymentRepository,
		donationRepository: donationRepository,
	}
}

func (ps *PaymentService) Create(input models.Payment) (models.Payment, error) {

	// get donation
	_, err := ps.donationRepository.GetByID(ps.db, input.DonationID)
	if err != nil {
		logrus.Printf("donation with ID : %v not found", input.DonationID)
		return models.Payment{}, errors.New("donation not found")
	}

	// create payment
	payment, err := ps.paymentRepository.Create(ps.db, input)
	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil

}

func (ps *PaymentService) GetByID(id uint) (models.Payment, error) {
	return ps.paymentRepository.GetByID(ps.db, id)
}
