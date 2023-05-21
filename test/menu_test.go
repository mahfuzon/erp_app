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

func setupMenuController(db *gorm.DB) *controllers.MenuController {
	menuRepository := repository.NewMenuRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	menuService := service.NewMenuService(menuRepository, categoryRepository)
	menuController := controllers.NewMenuController(menuService)
	return menuController
}

func truncateDataMenu(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE MENUS")
}

func createBulkExampleMenu(db *gorm.DB) {
	for i := 1; i <= 10; i++ {
		ingredient := models.Menu{Name: "menu " + strconv.Itoa(i), CategoryId: 1}
		db.Create(&ingredient)
	}
}

// test create success
func TestCreateSuccessMenu(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	createBulkExampleCategory(db)

	createRequestJson := `{
  "name" : "nasi goreng",
"category_id" : 1
}`

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu", menuController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu", strings.NewReader(createRequestJson))
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
func TestCreateFailMenuErrorValidation(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	category := models.Category{
		Id:   0,
		Name: "makanan",
	}

	err := db.Create(&category).Error
	assert.NoError(t, err)

	createRequestJson := `{
}`

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu", menuController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu", strings.NewReader(createRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 422, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test id not found
func TestCreateFailMenuErrorNotFoundCategory(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	category := models.Category{
		Id:   0,
		Name: "makanan",
	}

	err := db.Create(&category).Error
	assert.NoError(t, err)

	createRequestJson := `{
  "name" : "name",
"category_id" : 100
}`

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/menu", menuController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/menu", strings.NewReader(createRequestJson))
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

// test update success
func TestUpdateSuccessMenu(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	// create category
	category := models.Category{
		Id:   0,
		Name: "makanan",
	}
	err := db.Create(&category).Error
	assert.NoError(t, err)

	category = models.Category{
		Id:   0,
		Name: "minuman",
	}
	err = db.Create(&category).Error
	assert.NoError(t, err)
	// end

	// create menu
	menu := models.Menu{
		Id:         0,
		Name:       "before update",
		CategoryId: 1,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)
	// end

	createRequestJson := `{
  "name" : "jus pokat",
"category_id" : 2
}`

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.PUT("api/v1/menu/:id", menuController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/menu/1", strings.NewReader(createRequestJson))
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

// test get success
func TestGetSuccessMenu(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	// create category
	category := models.Category{
		Id:   0,
		Name: "makanan",
	}
	err := db.Create(&category).Error
	assert.NoError(t, err)

	category = models.Category{
		Id:   0,
		Name: "minuman",
	}
	err = db.Create(&category).Error
	assert.NoError(t, err)
	// end

	// create menu
	menu := models.Menu{
		Id:         0,
		Name:       "before update",
		CategoryId: 2,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)
	// end

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/menu/:id", menuController.Get)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/menu/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test get all with no query
func TestGetAllSuccessMenuWithNoQuery(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	createBulkExampleCategory(db)
	createBulkExampleIngredient(db)

	// create menu
	menu := models.Menu{
		Id:         0,
		Name:       "nasi goreng",
		CategoryId: 1,
	}
	err := db.Create(&menu).Error
	assert.NoError(t, err)

	menu = models.Menu{
		Id:         0,
		Name:       "jus pokat",
		CategoryId: 2,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)
	// end

	menu_ingredient_1 := models.MenuIngredient{
		MenuId:       1,
		IngredientId: 1,
		Qty:          "5 batang",
	}

	menu_ingredient_2 := models.MenuIngredient{
		MenuId:       1,
		IngredientId: 2,
		Qty:          "5 buah",
	}

	db.Create(&menu_ingredient_1)
	db.Create(&menu_ingredient_2)

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/menu", menuController.GetAll)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/menu", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test get all with query string
func TestGetAllSuccessMenuWithQuery(t *testing.T) {
	db := database.SetDbTest()
	truncateDataMenu(db)
	truncateDataCategory(db)

	// create category
	category := models.Category{
		Id:   0,
		Name: "makanan",
	}
	err := db.Create(&category).Error
	assert.NoError(t, err)

	category = models.Category{
		Id:   0,
		Name: "minuman",
	}
	err = db.Create(&category).Error
	assert.NoError(t, err)
	// end

	// create menu
	menu := models.Menu{
		Id:         0,
		Name:       "nasi goreng",
		CategoryId: 1,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)

	menu = models.Menu{
		Id:         0,
		Name:       "jus pokat",
		CategoryId: 2,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)
	// end

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.GET("api/v1/menu", menuController.GetAll)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/menu?name=jus", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}

// test delete success
func TestDeleteSuccessMenu(t *testing.T) {
	db := database.SetDbTest()

	db.Exec("TRUNCATE TABLE MENUS")
	db.Exec("TRUNCATE TABLE CATEGORIES")

	// create category
	category := models.Category{
		Id:   0,
		Name: "makanan",
	}
	err := db.Create(&category).Error
	assert.NoError(t, err)

	category = models.Category{
		Id:   0,
		Name: "minuman",
	}
	err = db.Create(&category).Error
	assert.NoError(t, err)
	// end

	// create menu
	menu := models.Menu{
		Id:         0,
		Name:       "nasi goreng",
		CategoryId: 1,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)

	menu = models.Menu{
		Id:         0,
		Name:       "jus pokat",
		CategoryId: 2,
	}
	err = db.Create(&menu).Error
	assert.NoError(t, err)
	// end

	menuController := setupMenuController(db)

	router := libraries.SetRouter()
	router.DELETE("api/v1/menu/:id", menuController.Delete)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/menu/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	assert.NoError(t, err)

	fmt.Println(data)
}
