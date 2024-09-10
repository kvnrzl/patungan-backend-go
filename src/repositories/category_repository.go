package repositories

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
}

func InitCategoryRepository() CategoryRepository {
	return CategoryRepository{}
}

func (cr *CategoryRepository) GetAll(tx *gorm.DB) ([]models.Category, error) {
	var categories []models.Category

	if err := tx.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepository) GetByID(tx *gorm.DB, id uint) (models.Category, error) {
	var category models.Category

	if err := tx.First(&category, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}
