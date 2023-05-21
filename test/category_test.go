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

func setupCategoryController(db *gorm.DB) *controllers.CategoryController {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	return controllers.NewCategoryController(categoryService)
}

func truncateDataCategory(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE CATEGORIES").Error
	if err != nil {
		panic(err.Error())
	}
}

func createBulkExampleCategory(db *gorm.DB) {
	for i := 1; i <= 10; i++ {
		category := models.Category{Name: "category " + strconv.Itoa(i)}
		db.Create(&category)
	}
}

// testing tanpa queryString
func TestGetAllSuccess(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/categories", categorycontroller.GetAll)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/categories", nil)
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

// testing dengan queryString
func TestGetAllWithStringQuery(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/categories", categorycontroller.GetAll)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/categories?name=category+2", nil)
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

// test jika data category tidak ditemukan
func TestGetIfEmptyData(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/categories/:id", categorycontroller.Get)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/categories/99", nil)
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

// test get data
func TestGetSuccess(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/categories/:id", categorycontroller.Get)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/categories/10", nil)
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

// test create success
func TestCreateSuccess(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)

	createRequestJson := `{
  "name" : "category 10"
}`

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/categories", categorycontroller.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/categories", strings.NewReader(createRequestJson))
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

// test jika yang fikirimkan string kosong
func TestCreateFailValidationFormEmptyString(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)

	createRequestJson := `{
  "name" : ""
}`

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/categories", categorycontroller.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/categories", strings.NewReader(createRequestJson))
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

// test jika tidak mengirimkan payload
func TestCreateFailValidationNullField(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/categories", categorycontroller.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/categories", nil)
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

// test update success
func TestUpdateSuccess(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	updateRequestJson := `{
  "name" : "category 99"
}`

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/categories/:id", categorycontroller.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/categories/1", strings.NewReader(updateRequestJson))
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

func TestUpdateFailInvalidParameter(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	updateRequestJson := `{
  "name" : "category 99"
}`

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/categories/:id", categorycontroller.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/categories/ss", strings.NewReader(updateRequestJson))
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

func TestUpdateFailNotParameter(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	updateRequestJson := `{
  "name" : "category 99"
}`

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/categories/:id", categorycontroller.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/categories", strings.NewReader(updateRequestJson))
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

func TestUpdateFailNotFoundDataForUpdate(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	updateRequestJson := `{
  "name" : "category 99"
}`

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/categories/:id", categorycontroller.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/categories/99", strings.NewReader(updateRequestJson))
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

func TestDeleteSuccess(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/categories/:id", categorycontroller.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/categories/10", nil)
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

func TestDeleteFailNotFoundData(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/categories/:id", categorycontroller.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/categories/11", nil)
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

func TestDeleteFailNotExistsParameter(t *testing.T) {
	db := database.SetDbTest()
	truncateDataCategory(db)
	createBulkExampleCategory(db)

	categorycontroller := setupCategoryController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/categories/:id", categorycontroller.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/categories/", nil)
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
