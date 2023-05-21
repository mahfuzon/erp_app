package repository

import (
	"fmt"
	"github.com/erp_app/models"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	Create(recipe models.MenuIngredient) (models.MenuIngredient, error)
	Update(recipe models.MenuIngredient) (models.MenuIngredient, error)
	Find(id int) (models.MenuIngredient, error)
	All() ([]models.MenuIngredient, error)
	Delete(recipe models.MenuIngredient) error
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

func (recipeRepository *recipeRepository) Create(recipe models.MenuIngredient) (models.MenuIngredient, error) {
	err := recipeRepository.db.Create(&recipe).Error
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (recipeRepository *recipeRepository) Update(recipe models.MenuIngredient) (models.MenuIngredient, error) {
	err := recipeRepository.db.Save(&recipe).Error
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (recipeRepository *recipeRepository) Find(id int) (models.MenuIngredient, error) {
	fmt.Println(id)
	recipe := models.MenuIngredient{}
	err := recipeRepository.db.Preload("Ingredient").First(&recipe, id).Error
	if err != nil {
		return recipe, err
	}

	fmt.Println(recipe)

	return recipe, nil
}

func (recipeRepository *recipeRepository) All() ([]models.MenuIngredient, error) {
	var listRecipe []models.MenuIngredient

	err := recipeRepository.db.Preload("Category").Find(&listRecipe).Error

	if err != nil {
		return listRecipe, err
	}

	return listRecipe, nil
}

func (recipeRepository *recipeRepository) Delete(recipe models.MenuIngredient) error {
	err := recipeRepository.db.Delete(&recipe).Error
	if err != nil {
		return err
	}

	return nil
}
