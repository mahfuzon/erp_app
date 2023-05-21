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

type RecipeController struct {
	recipeService service.RecipeService
}

func NewRecipeController(recipeService service.RecipeService) *RecipeController {
	return &RecipeController{recipeService: recipeService}
}

func (recipeController *RecipeController) Add(ctx echo.Context) error {
	createRecipeRequest := request.CreateRecipeRequest{}

	err := ctx.Bind(&createRecipeRequest)
	if err != nil {
		fmt.Println("error binding")
		apiResponse := response.NewApiResponse("error", "failed create menu recipe", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&createRecipeRequest)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed create menu recipes", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	menuResponse, err := recipeController.recipeService.Create(createRecipeRequest)
	if err != nil {
		fmt.Println("error service")
		apiResponse := response.NewApiResponse("error", "failed create menu recipes", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success create menu recipes", menuResponse)
	return ctx.JSON(201, apiResponse)
}

func (recipeController *RecipeController) Update(ctx echo.Context) error {
	req := request.UpdateRecipeRequest{}

	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println("error binding")
		apiResponse := response.NewApiResponse("error", "failed update menu recipe", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&req)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed update menu recipes", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	menuResponse, err := recipeController.recipeService.Update(req)
	if err != nil {
		fmt.Println("error service")
		apiResponse := response.NewApiResponse("error", "failed update menu recipes", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success create menu recipes", menuResponse)
	return ctx.JSON(201, apiResponse)
}

func (recipeController *RecipeController) Delete(ctx echo.Context) error {
	req := request.DeleteRecipeRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println("error binding")
		apiResponse := response.NewApiResponse("error", "failed delete menu recipe", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&req)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed delete menu recipes", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	err = recipeController.recipeService.Delete(req)
	if err != nil {
		fmt.Println("error service")
		apiResponse := response.NewApiResponse("error", "failed delete menu recipes", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success delete menu recipes", nil)
	return ctx.JSON(201, apiResponse)
}
