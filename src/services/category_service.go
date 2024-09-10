package services

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/repositories"
	"gorm.io/gorm"
)

type CategoryService struct {
	db         *gorm.DB
	repository repositories.CategoryRepository
}

func InitCategoryService(db *gorm.DB, repository repositories.CategoryRepository) CategoryService {
	return CategoryService{
		db:         db,
		repository: repository,
	}
}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	return cs.repository.GetAll(cs.db)
}

func (cs *CategoryService) GetByID(id uint) (models.Category, error) {
	return cs.repository.GetByID(cs.db, id)
}
