package service

import (
	"fmt"
	"github.com/erp_app/models"
	"github.com/erp_app/repository"
	"github.com/erp_app/request"
)

type RecipeService interface {
	Create(createRecipeRequest request.CreateRecipeRequest) (models.MenuIngredient, error)
	Update(recipeRequest request.UpdateRecipeRequest) (models.MenuIngredient, error)
	Delete(recipeRequest request.DeleteRecipeRequest) error
}

type recipeService struct {
	recipeRepository     repository.RecipeRepository
	menuRepository       repository.MenuRepository
	ingredientRepository repository.IngredientRepository
}

func NewRecipeService(recipeRepository repository.RecipeRepository, menuRepository repository.MenuRepository, ingredientRepository repository.IngredientRepository) RecipeService {
	return &recipeService{
		recipeRepository:     recipeRepository,
		menuRepository:       menuRepository,
		ingredientRepository: ingredientRepository,
	}
}

func (recipeService *recipeService) Create(createRecipeRequest request.CreateRecipeRequest) (models.MenuIngredient, error) {
	recipe := models.MenuIngredient{}
	menu, err := recipeService.menuRepository.Find(createRecipeRequest.MenuId)
	if err != nil {
		fmt.Println("menu not found")
		return recipe, err
	}

	ingredient, err := recipeService.ingredientRepository.Find(createRecipeRequest.IngredientId)
	if err != nil {
		fmt.Println("ingredient not found")
		return recipe, err
	}

	recipe.MenuId = menu.Id
	recipe.IngredientId = ingredient.Id
	recipe.Qty = createRecipeRequest.Qty
	recipe.Ingredient = ingredient

	recipe, err = recipeService.recipeRepository.Create(recipe)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (recipeService *recipeService) Update(recipeRequest request.UpdateRecipeRequest) (models.MenuIngredient, error) {
	fmt.Println(recipeRequest)
	recipe, err := recipeService.recipeRepository.Find(recipeRequest.Id)
	if err != nil {
		fmt.Println("error recipe find")
		return recipe, err
	}

	ingredient, err := recipeService.ingredientRepository.Find(recipeRequest.IngredientId)
	if err != nil {
		return recipe, err
	}

	menu, err := recipeService.menuRepository.Find(recipeRequest.MenuId)
	if err != nil {
		return recipe, err
	}

	recipe.MenuId = menu.Id
	recipe.IngredientId = ingredient.Id
	recipe.Qty = recipeRequest.Qty
	recipe.Ingredient = ingredient

	recipe, err = recipeService.recipeRepository.Update(recipe)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (recipeService *recipeService) Delete(recipeRequest request.DeleteRecipeRequest) error {
	recipe, err := recipeService.recipeRepository.Find(recipeRequest.Id)
	if err != nil {
		fmt.Println("error recipe find")
		return err
	}

	err = recipeService.recipeRepository.Delete(recipe)
	if err != nil {
		return err
	}

	return nil
}
