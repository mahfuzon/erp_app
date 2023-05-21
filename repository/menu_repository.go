package repository

import (
	"github.com/erp_app/models"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu models.Menu) (models.Menu, error)
	Update(menu models.Menu) (models.Menu, error)
	Find(id int) (models.Menu, error)
	All(name string) ([]models.Menu, error)
	Delete(ingredient models.Menu) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

func (menuRepository *menuRepository) Create(menu models.Menu) (models.Menu, error) {
	err := menuRepository.db.Create(&menu).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (menuRepository *menuRepository) Update(menu models.Menu) (models.Menu, error) {
	err := menuRepository.db.Save(&menu).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (menuRepository *menuRepository) Find(id int) (models.Menu, error) {
	menu := models.Menu{}
	err := menuRepository.db.Preload("Category").Preload("Ingredients").Preload("Ingredients.Ingredient").First(&menu, id).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (menuRepository *menuRepository) All(name string) ([]models.Menu, error) {
	var listMenu []models.Menu
	query := menuRepository.db

	if name != "" {
		query = query.Where("name Like ?", "%"+name+"%")
	}

	err := query.Preload("Category").Preload("Ingredients").Preload("Ingredients.Ingredient").Find(&listMenu).Error

	if err != nil {
		return listMenu, err
	}

	return listMenu, nil
}

func (menuRepository *menuRepository) Delete(menu models.Menu) error {
	err := menuRepository.db.Delete(&menu).Error
	if err != nil {
		return err
	}

	return nil
}
