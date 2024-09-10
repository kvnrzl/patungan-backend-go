package repositories

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func InitUserRepository() UserRepository {
	return UserRepository{}
}

func (cr *UserRepository) GetAll(tx *gorm.DB) ([]models.User, error) {
	var users []models.User

	if err := tx.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (cr *UserRepository) GetByID(tx *gorm.DB, id uint) (models.User, error) {
	var user models.User

	if err := tx.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (cr *UserRepository) GetByEmail(tx *gorm.DB, email string) (models.User, error) {
	var user models.User

	if err := tx.First(&user, "email = ?", email).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (cr *UserRepository) Create(tx *gorm.DB, userInput models.User) (models.User, error) {

	var user models.User

	result := tx.Create(&userInput)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := result.Last(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (cr *UserRepository) Update(tx *gorm.DB, userInput models.User, id uint) (models.User, error) {
	user, err := cr.GetByID(tx, id)

	if err != nil {
		return models.User{}, err
	}

	//user.Email = userInput.Email (cannot be changed)
	user.Name = userInput.Name
	user.Phone = userInput.Phone
	if userInput.Password != "" {
		user.Password = userInput.Password
	}

	// check if is_verified exist or not
	if userInput.IsVerified {
		user.IsVerified = userInput.IsVerified
	}

	if err = tx.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (cr *UserRepository) Delete(tx *gorm.DB, id uint) error {
	user, err := cr.GetByID(tx, id)

	if err != nil {
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
