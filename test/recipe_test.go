package test

import (
	"encoding/json"
	"fmt"
	"github.com/erp_app/controllers"
	"github.com/erp_app/database"
	"github.com/erp_app/libraries"
	"github.com/erp_app/models"
	"github.com/erp_app/repository"
	"github.com/erp_app/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func truncateDataRecipes(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE RECIPES")
}

func setupRecipeController(db *gorm.DB) *controllers.RecipeController {
	recipeRepository := repository.NewRecipeRepository(db)
	menuRepository := repository.NewMenuRepository(db)
	ingredientRepository := repository.NewIngredientRepository(db)
	recipeService := service.NewRecipeService(recipeRepository, menuRepository, ingredientRepository)
	recipeController := controllers.NewRecipeController(recipeService)
	return recipeController
}

// test add success
func TestAddSuccess(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu/:menu_id/recipes", recipeController.Add)

	createRequestJson := `{
  "ingredient_id" : 1,
  "description" : "siung",
"qty" : "5"
}`

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu/1/recipes", strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 201, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test validation
func TestAddFailValidation(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu/:menu_id/recipes", recipeController.Add)

	createRequestJson := `{
}`

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu//recipes", strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 422, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test menu not found
func TestAddFailMenuNotFound(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu/:menu_id/recipes", recipeController.Add)

	createRequestJson := `{
  "ingredient_id" : 1,
"qty" : "5"
}`

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu/200/recipes", strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 500, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test ingredient not found
func TestAddFailIngredientNotFound(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu/:menu_id/recipes", recipeController.Add)

	createRequestJson := `{
  "menu_id" : 1,
  "ingredient_id" : 100,
  "description" : "siung",
"qty" : "5"
}`

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu/1/recipes", strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 500, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test update success
func TestUpdateSuccessRecipe(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	recipe := models.MenuIngredient{
		MenuId:       2,
		IngredientId: 2,
		Qty:          "2 buah",
	}

	err := db.Create(&recipe).Error
	assert.NoError(t, err)

	router := libraries.SetRouter()
	router.PUT("api/v1/menu/:menu_id/recipes/:id", recipeController.Update)

	createRequestJson := `{
  "ingredient_id" : 1,
"qty" : "5"
}`

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/menu/"+strconv.Itoa(recipe.MenuId)+"/recipes/"+strconv.Itoa(recipe.Id), strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 201, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test id not found
func TestUpdateFailRecipeIdNotFound(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	recipe := models.MenuIngredient{
		MenuId:       2,
		IngredientId: 2,
		Qty:          "2 buah",
	}

	err := db.Create(&recipe).Error
	assert.NoError(t, err)

	router := libraries.SetRouter()
	router.PUT("api/v1/menu/:menu_id/recipes/:id", recipeController.Update)

	createRequestJson := `{
  "ingredient_id" : 1,
"qty" : "5"
}`

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/menu/"+strconv.Itoa(recipe.MenuId)+"/recipes/100", strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test delete success
func TestDeleteSuccessRecipe(t *testing.T) {
	db := database.SetDbTest()
	truncateDataRecipes(db)
	truncateDataCategory(db)
	truncateDataMenu(db)
	truncateDataIngredient(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)
	createBulkExampleMenu(db)

	recipeController := setupRecipeController(db)

	recipe := models.MenuIngredient{
		MenuId:       2,
		IngredientId: 2,
		Qty:          "2 buah",
	}

	err := db.Create(&recipe).Error
	assert.NoError(t, err)

	router := libraries.SetRouter()
	router.DELETE("api/v1/menu/:menu_id/recipes/:id", recipeController.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/menu/"+strconv.Itoa(recipe.MenuId)+"/recipes/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 201, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}
