package services

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/repositories"
	"errors"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

// var MidtransServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
var MidtransServerKey = "SB-Mid-server-f93hihpd4YtVdOOYsp0AJKgi"

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

func (ps *PaymentService) CreateTransaction(input models.Payment) (models.Payment, *snap.Response, error) {

	// get donation
	_, err := ps.donationRepository.GetByID(ps.db, input.DonationID)
	if err != nil {
		logrus.Printf("donation with ID : %v not found", input.DonationID)
		return models.Payment{}, nil, errors.New("donation not found")
	}

	// get user donation
	user, _ := ps.donationRepository.GetUserDonation(ps.db, input.DonationID)

	// ====================================================================
	// MIDTRANS API

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(MidtransServerKey, midtrans.Sandbox)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(input.DonationID)),
			GrossAmt: int64(input.Amount),
		},
		//CreditCard: &snap.CreditCardDetails{
		//	Secure: true,
		//},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)
	fmt.Println("snapResp :", snapResp)

	if snapResp.ErrorMessages != nil {
		logrus.Println("snapResp.ErrorMessages: ", snapResp.ErrorMessages)
		return models.Payment{}, nil, errors.New("error create payment")
	}

	// ====================================================================

	// create payment
	payment, err := ps.paymentRepository.Create(ps.db, input)
	if err != nil {
		return models.Payment{}, nil, err
	}

	return payment, snapResp, nil

}

func (ps *PaymentService) GetByID(id uint) (models.Payment, error) {
	return ps.paymentRepository.GetByID(ps.db, id)
}
