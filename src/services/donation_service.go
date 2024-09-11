package services

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/repositories"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DonationService struct {
	db                 *gorm.DB
	donationRepository repositories.DonationRepository
	userRepository     repositories.UserRepository
}

func InitDonationService(db *gorm.DB, donationRepository repositories.DonationRepository, userRepository repositories.UserRepository) DonationService {
	return DonationService{
		db:                 db,
		donationRepository: donationRepository,
		userRepository:     userRepository,
	}
}

func (ds *DonationService) CreateDonationGuestUser(donation models.Donation) (models.Donation, error) {

	// check user exist
	user, err := ds.userRepository.GetByEmail(ds.db, donation.User.Email)
	if err != nil {
		fmt.Println("User not found, creating new user")
		// if user not exist, create new user
		user, err = ds.userRepository.Create(ds.db, donation.User)
		if err != nil {
			return models.Donation{}, errors.New("failed to create new user")
		}
	}

	// Ensure user ID is valid
	if user.ID == 0 {
		return models.Donation{}, errors.New("invalid user ID")
	}

	// Set user id to donation
	donation.UserID = user.ID
	fmt.Println("donation cuyyy: ", donation)
	fmt.Println("donation.UserID: ", donation.UserID)

	// then create donation
	return ds.donationRepository.Create(ds.db, donation)
}

//func (ds *DonationService) CreateDonationGuestUser(donation models.Donation) (models.Donation, error) {
//
//	// Check if user exists with the provided email
//	user, err := ds.userRepository.GetByEmail(ds.db, donation.User.Email)
//	if err == nil {
//		// User found, set donation.UserID
//		donation.UserID = user.ID
//		fmt.Println("Donation with existing user:", donation)
//	} else if errors.Is(err, sql.ErrNoRows) { // Check for specific "no rows" error
//		// User not found, attempt to create new user
//		newUser, err := ds.userRepository.Create(ds.db, donation.User)
//		if err != nil {
//			return models.Donation{}, errors.New("failed to create new user")
//		}
//		donation.UserID = newUser.ID
//		fmt.Println("Donation with newly created user:", donation)
//	} else {
//		// Handle other potential errors from GetByEmail
//		return models.Donation{}, errors.New("error retrieving user")
//	}
//
//	// Create donation with validated UserID
//	return ds.donationRepository.Create(ds.db, donation)
//}

func (ds *DonationService) CreateDonationRegisteredUser(donation models.Donation) (models.Donation, error) {

	// jika ada maka langsung create donation
	return ds.donationRepository.Create(ds.db, donation)

}
