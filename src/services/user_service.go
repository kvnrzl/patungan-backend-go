package services

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/pkg/jwt"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/src/repositories"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db             *gorm.DB
	userRepository repositories.UserRepository
	validate       *validator.Validate
}

func InitUserService(db *gorm.DB, repository repositories.UserRepository) UserService {
	return UserService{
		db:             db,
		userRepository: repository,
	}
}

func (us *UserService) Login(user models.User) (string, error) {

	// validate first
	var userLoginRequest dto.UserLoginRequest
	userLoginRequest.ToEntity()

	// get user from db
	newUser, err := us.userRepository.GetByEmail(us.db, user.Email)
	if err != nil {
		logrus.Println("error get user by email :", err)
		return "", errors.New("invalid email or password")
	}
	fmt.Println(newUser)

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password))
	if err != nil {
		logrus.Println("error compare password :", err)
		return "", errors.New("invalid email or password")
	}

	// generate jwt
	token, err := jwt.GenerateToken(newUser, 1)
	if err != nil {
		logrus.Println("error generate token :", err)
		return "", errors.New("error generate token")
	}

	return token, nil
}

func (us *UserService) Register(user models.User) (models.User, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Println("error hash password :", err)
		return models.User{}, errors.New("error hash password")
	}

	user.Password = string(hashedPassword)

	// create user
	newUser, err := us.userRepository.Create(us.db, user)
	if err != nil {
		logrus.Println("error create user :", err)
		return models.User{}, errors.New("error create user")
	}

	return newUser, nil
}
