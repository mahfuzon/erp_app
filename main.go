package main

import (
	"github.com/erp_app/controllers"
	"github.com/erp_app/database"
	"github.com/erp_app/libraries"
	"github.com/erp_app/repository"
	"github.com/erp_app/service"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db := database.SetDb()
	router := libraries.SetRouter()

	apiV1 := router.Group("/api/v1")

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	apiV1Category := apiV1.Group("/category")
	apiV1Category.GET("", categoryController.GetAll)
	apiV1Category.GET("/:id", categoryController.Get)
	apiV1Category.POST("", categoryController.Create)
	apiV1Category.PUT("/:id", categoryController.Update)
	apiV1Category.DELETE("/:id", categoryController.Delete)

	ingredientRepository := repository.NewIngredientRepository(db)
	IngredientService := service.NewIngredientService(ingredientRepository)
	ingredientController := controllers.NewIngredientController(IngredientService)

	apiV1Ingredient := apiV1.Group("/ingredient")
	apiV1Ingredient.GET("", ingredientController.GetAll)
	apiV1Ingredient.GET("/:id", ingredientController.Get)
	apiV1Ingredient.POST("", ingredientController.Create)
	apiV1Ingredient.PUT("/:id", ingredientController.Update)
	apiV1Ingredient.DELETE("/:id", ingredientController.Delete)

	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository, categoryRepository)
	menuController := controllers.NewMenuController(menuService)
	recipeRepository := repository.NewRecipeRepository(db)
	recipeService := service.NewRecipeService(recipeRepository, menuRepository, ingredientRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	apiV1Menu := apiV1.Group("/menu")
	apiV1Menu.GET("", menuController.GetAll)
	apiV1Menu.GET("/:id", menuController.Get)
	apiV1Menu.POST("", menuController.Create)
	apiV1Menu.PUT("/:id", menuController.Update)
	apiV1Menu.DELETE("/:id", menuController.Delete)
	apiV1Menu.POST("/:menu_id/recipe/", recipeController.Add)
	apiV1Menu.PUT("/:menu_id/recipe/:id", recipeController.Update)
	apiV1Menu.DELETE("/:menu_id/recipe/:id", recipeController.Delete)

	router.Logger.Fatal(router.Start(":8000"))
}
