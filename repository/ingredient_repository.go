package repository

import (
	"github.com/erp_app/models"
	"gorm.io/gorm"
)

type IngredientRepository interface {
	All(name string) ([]models.Ingredient, error)
	Find(id int) (models.Ingredient, error)
	Create(ingredient models.Ingredient) (models.Ingredient, error)
	Update(ingredient models.Ingredient) (models.Ingredient, error)
	Delete(ingredient models.Ingredient) error
}

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IngredientRepository {
	return &ingredientRepository{
		db: db,
	}
}

func (ingredientRepository *ingredientRepository) All(name string) ([]models.Ingredient, error) {
	var listIngredient []models.Ingredient
	query := ingredientRepository.db

	if name != "" {
		query = query.Where("name Like ?", "%"+name+"%")
	}

	err := query.Find(&listIngredient).Error

	if err != nil {
		return listIngredient, err
	}

	return listIngredient, nil
}

func (ingredientRepository *ingredientRepository) Find(id int) (models.Ingredient, error) {
	ingredient := models.Ingredient{}
	err := ingredientRepository.db.First(&ingredient, id).Error
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

func (ingredientRepository *ingredientRepository) Create(ingredient models.Ingredient) (models.Ingredient, error) {
	err := ingredientRepository.db.Create(&ingredient).Error
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

func (ingredientRepository *ingredientRepository) Update(ingredient models.Ingredient) (models.Ingredient, error) {
	err := ingredientRepository.db.Save(&ingredient).Error
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

func (ingredientRepository *ingredientRepository) Delete(ingredient models.Ingredient) error {
	err := ingredientRepository.db.Delete(&ingredient).Error
	if err != nil {
		return err
	}

	return nil
}
