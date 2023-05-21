package controllers

import (
	"fmt"
	"github.com/erp_app/helper"
	"github.com/erp_app/request"
	"github.com/erp_app/response"
	"github.com/erp_app/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type IngredientController struct {
	IngredientService service.IngredientService
}

func NewIngredientController(ingredientService service.IngredientService) *IngredientController {
	return &IngredientController{IngredientService: ingredientService}
}

func (ingredientController *IngredientController) GetAll(ctx echo.Context) error {
	getAllRequestIngredient := request.GetAllRequestIngredient{}
	err := ctx.Bind(&getAllRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get all ingredient", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&getAllRequestIngredient)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed get all ingredient", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	listIngredientResponse, err := ingredientController.IngredientService.GetAll(getAllRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get all ingredient", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success get all ingredient", listIngredientResponse)
	return ctx.JSON(200, apiResponse)
}

func (ingredientController *IngredientController) Get(ctx echo.Context) error {
	getDetailRequestIngredient := request.GetDetailRequestIngredient{}
	err := ctx.Bind(&getDetailRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail ingredient", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&getDetailRequestIngredient)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed get detail ingredient", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	ingredientResponse, err := ingredientController.IngredientService.Get(getDetailRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail ingredient", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success get detail ingredient", ingredientResponse)
	return ctx.JSON(200, apiResponse)
}

func (ingredientController *IngredientController) Delete(ctx echo.Context) error {
	deleteRequestIngredient := request.DeleteRequestIngredient{}
	err := ctx.Bind(&deleteRequestIngredient)
	if err != nil {
		fmt.Println("error binding")
		apiResponse := response.NewApiResponse("error", "failed delete ingredient", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&deleteRequestIngredient)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed delete ingredient", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	err = ingredientController.IngredientService.Delete(deleteRequestIngredient)
	if err != nil {
		fmt.Println("error service")
		apiResponse := response.NewApiResponse("error", "failed delete ingredient", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success delete ingredient", nil)
	return ctx.JSON(200, apiResponse)
}

func (ingredientController *IngredientController) Create(ctx echo.Context) error {
	createRequestIngredient := request.CreateRequestIngredient{}
	err := ctx.Bind(&createRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed create ingredient", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&createRequestIngredient)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed create ingredient", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	ingredientResponse, err := ingredientController.IngredientService.Create(createRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed create ingredient", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success create ingredient", ingredientResponse)
	return ctx.JSON(201, apiResponse)
}

func (ingredientController *IngredientController) Update(ctx echo.Context) error {
	updateRequestIngredient := request.UpdateRequestIngredient{}
	err := ctx.Bind(&updateRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed update ingredient", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&updateRequestIngredient)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed update ingredient", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	ingredientResponse, err := ingredientController.IngredientService.Update(updateRequestIngredient)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed update ingredient", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success update ingredient", ingredientResponse)
	return ctx.JSON(201, apiResponse)
}
