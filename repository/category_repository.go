package repository

import (
	"github.com/erp_app/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	All(name string) ([]models.Category, error)
	Find(id int) (models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(category models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (categoryRepository *categoryRepository) All(name string) ([]models.Category, error) {
	var listCategories []models.Category
	query := categoryRepository.db

	if name != "" {
		query = query.Where("name Like ?", "%"+name+"%")
	}

	err := query.Find(&listCategories).Error

	if err != nil {
		return listCategories, err
	}

	return listCategories, nil
}

func (categoryRepository *categoryRepository) Find(id int) (models.Category, error) {
	category := models.Category{}
	err := categoryRepository.db.First(&category, id).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (categoryRepository *categoryRepository) Create(category models.Category) (models.Category, error) {
	err := categoryRepository.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (categoryRepository *categoryRepository) Update(category models.Category) (models.Category, error) {
	err := categoryRepository.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (categoryRepository *categoryRepository) Delete(category models.Category) error {
	err := categoryRepository.db.Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}
