package controllers

import (
	"github.com/erp_app/helper"
	"github.com/erp_app/request"
	"github.com/erp_app/response"
	"github.com/erp_app/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (categoryController *CategoryController) GetAll(ctx echo.Context) error {
	getAllRequestCategory := request.GetAllRequestCategory{}
	err := ctx.Bind(&getAllRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get all category", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&getAllRequestCategory)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed get all category", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	listCategoryResponse, err := categoryController.CategoryService.GetAll(getAllRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get all category", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success get all category", listCategoryResponse)
	return ctx.JSON(200, apiResponse)
}

func (categoryController *CategoryController) Get(ctx echo.Context) error {
	getDetailRequestCategory := request.GetDetailRequestCategory{}
	err := ctx.Bind(&getDetailRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail category", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&getDetailRequestCategory)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed get detail category", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	categoryResponse, err := categoryController.CategoryService.Get(getDetailRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail category", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success get detail category", categoryResponse)
	return ctx.JSON(200, apiResponse)
}

func (categoryController *CategoryController) Delete(ctx echo.Context) error {
	deleteRequestCategory := request.DeleteRequestCategory{}
	err := ctx.Bind(&deleteRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed delete category", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&deleteRequestCategory)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed delete category", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	err = categoryController.CategoryService.Delete(deleteRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed delete category", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success delete category", nil)
	return ctx.JSON(200, apiResponse)
}

func (categoryController *CategoryController) Create(ctx echo.Context) error {
	createRequestCategory := request.CreateRequestCategory{}
	err := ctx.Bind(&createRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed create category", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&createRequestCategory)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed create category", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	categoryResponse, err := categoryController.CategoryService.Create(createRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed create category", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success create category", categoryResponse)
	return ctx.JSON(201, apiResponse)
}

func (categoryController *CategoryController) Update(ctx echo.Context) error {
	updateRequestCategory := request.UpdateRequestCategory{}
	err := ctx.Bind(&updateRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed update category", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&updateRequestCategory)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed update category", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	categoryResponse, err := categoryController.CategoryService.Update(updateRequestCategory)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed update category", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success update category", categoryResponse)
	return ctx.JSON(201, apiResponse)
}
