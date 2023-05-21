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

func setupIngredientController(db *gorm.DB) *controllers.IngredientController {
	ingredientRepository := repository.NewIngredientRepository(db)
	ingredientService := service.NewIngredientService(ingredientRepository)
	return controllers.NewIngredientController(ingredientService)
}

func truncateDataIngredient(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE INGREDIENTS").Error
	if err != nil {
		panic(err.Error())
	}
}

func createBulkExampleIngredient(db *gorm.DB) {
	for i := 1; i <= 10; i++ {
		ingredient := models.Ingredient{Name: "ingredient " + strconv.Itoa(i)}
		db.Create(&ingredient)
	}
}

// test get all ingredient
func TestGetAllWithDataIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/ingredient", ingredientController.GetAll)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/ingredient", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test get all with query
func TestGetAllWithStringQueryIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/ingredient", ingredientController.GetAll)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/ingredient?name=ingredient+2", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestGetIfEmptyDataIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/ingredient/:id", ingredientController.Get)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/ingredient/99", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestGetIfExistsDataIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/ingredient/:id", ingredientController.Get)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/ingredient/10", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestCreateSuccessIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)

	createRequestJson := `{
  "name" : "ingredient 10"
}`

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/ingredient", ingredientController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/ingredient", strings.NewReader(createRequestJson))
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

func TestCreateFailValidationFormEmptyStringIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)

	createRequestJson := `{
  "name" : ""
}`

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/ingredient", ingredientController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/ingredient", strings.NewReader(createRequestJson))
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

func TestCreateFailValidationNullFieldIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/ingredient", ingredientController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/ingredient", nil)
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

func TestUpdateSuccessIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	updateRequestJson := `{
  "name" : "ingredient 99"
}`

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/ingredient/:id", ingredientController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/ingredient/1", strings.NewReader(updateRequestJson))
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

func TestUpdateFailInvalidParameterIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	updateRequestJson := `{
  "name" : "ingredient 99"
}`

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/ingredient/:id", ingredientController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/ingredient/ss", strings.NewReader(updateRequestJson))
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

func TestUpdateFailNotParameterIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	updateRequestJson := `{
  "name" : "ingredient 99"
}`

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/ingredient/:id", ingredientController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/ingredient", strings.NewReader(updateRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 405, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestUpdateFailNotFounDataForUpdateIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	updateRequestJson := `{
  "name" : "ingredient 99"
}`

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/ingredient/:id", ingredientController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/ingredient/99", strings.NewReader(updateRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestDeleteSuccessIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/ingredient/:id", ingredientController.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/ingredient/10", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestDeleteFailNotFoundDataIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/ingredient/:id", ingredientController.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/ingredient/11", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

func TestDeleteFailNotExistsParameterIngredient(t *testing.T) {
	db := database.SetDbTest()
	truncateDataIngredient(db)
	createBulkExampleIngredient(db)

	ingredientController := setupIngredientController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/ingredient/:id", ingredientController.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/ingredient/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 405, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err := json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}
