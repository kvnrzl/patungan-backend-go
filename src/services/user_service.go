package services

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/dto"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/pkg/jwt"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/src/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db             *gorm.DB
	rdb            *redis.Client
	userRepository repositories.UserRepository
}

func InitUserService(db *gorm.DB, rdb *redis.Client, repository repositories.UserRepository) UserService {
	return UserService{
		db:             db,
		rdb:            rdb,
		userRepository: repository,
	}
}

func (us *UserService) Login(ctx context.Context, user models.User) (string, error) {

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
	uuidStr := uuid.New().String()
	token, err := jwt.GenerateToken(newUser, uuidStr, 1)
	if err != nil {
		logrus.Println("error generate token :", err)
		return "", errors.New("error generate token")
	}

	// set to redis
	err = us.rdb.Set(ctx, newUser.Email, uuidStr, 0).Err()
	if err != nil {
		logrus.Println("error set token to redis :", err)
		return "", errors.New("error set token to redis")
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

func (us *UserService) Logout(ctx context.Context, email string) error {
	// delete token from redis
	err := us.rdb.Del(ctx, email).Err()
	if err != nil {
		logrus.Println("error delete token from redis :", err)
		return errors.New("error delete token from redis")
	}

	return nil
}
