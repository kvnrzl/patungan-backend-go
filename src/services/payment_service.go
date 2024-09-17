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
	campaignRepository repositories.CampaignRepository
}

func InitPaymentService(db *gorm.DB, paymentRepository repositories.PaymentRepository, donationRepository repositories.DonationRepository, campaignRepository repositories.CampaignRepository) PaymentService {
	return PaymentService{
		db:                 db,
		paymentRepository:  paymentRepository,
		donationRepository: donationRepository,
		campaignRepository: campaignRepository,
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

	// create payment
	payment, err := ps.paymentRepository.Create(ps.db, input)
	if err != nil {
		return models.Payment{}, nil, err
	}
	// ====================================================================
	// MIDTRANS API

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(MidtransServerKey, midtrans.Sandbox)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(payment.ID)),
			GrossAmt: int64(input.Amount),
		},
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

	return payment, snapResp, nil

}

func (ps *PaymentService) UpdatePaymentStatus(payment models.Payment) error {

	// update payment status
	tx := ps.db.Begin()

	payment, err := ps.paymentRepository.Update(tx, payment)
	if err != nil {
		logrus.Println("error update payment status: ", err)
		tx.Rollback()
		return err
	}
	fmt.Println("payment updated: ", payment)

	// if status 'settlement' then update donation status
	if payment.PaymentStatus == "settlement" {
		donation, err := ps.donationRepository.GetByID(tx, payment.DonationID)
		if err != nil {
			logrus.Println("error get donation: ", err)
			tx.Rollback()
			return err
		}

		//update the campaign collected
		campaign, err := ps.campaignRepository.GetByID(tx, donation.CampaignID)
		if err != nil {
			logrus.Println("error get campaign: ", err)
			tx.Rollback()
			return err
		}

		campaign.Collected = campaign.Collected + payment.Amount
		_, err = ps.campaignRepository.UpdateCollected(tx, campaign)
		if err != nil {
			logrus.Println("error update campaign collected amount: ", err)
			tx.Rollback()
			return err
		}
		logrus.Println("campaign collected updated: ", campaign.Collected)
	}

	tx.Commit()

	fmt.Println("payment: ", payment)

	return nil

}

func (ps *PaymentService) GetByID(id uint) (models.Payment, error) {
	return ps.paymentRepository.GetByID(ps.db, id)
}
