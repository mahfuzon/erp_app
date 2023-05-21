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

type MenuController struct {
	menuService service.MenuService
}

func NewMenuController(menuService service.MenuService) *MenuController {
	return &MenuController{
		menuService: menuService,
	}
}

func (menuController *MenuController) Delete(ctx echo.Context) error {
	deleteMenuRequest := request.DeleteMenuRequest{}
	err := ctx.Bind(&deleteMenuRequest)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed delete menu", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&deleteMenuRequest)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed delete menu", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	err = menuController.menuService.Delete(deleteMenuRequest)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed delete menu", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success delete menu", nil)
	return ctx.JSON(200, apiResponse)
}

func (menuController *MenuController) Create(ctx echo.Context) error {
	createMenuRequest := request.CreateMenuRequest{}
	err := ctx.Bind(&createMenuRequest)
	if err != nil {
		fmt.Println("error binding")
		apiResponse := response.NewApiResponse("error", "failed create menu", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&createMenuRequest)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed create menu", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	menuResponse, err := menuController.menuService.Create(createMenuRequest)
	if err != nil {
		fmt.Println("error service")
		apiResponse := response.NewApiResponse("error", "failed create menu", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success create menu", menuResponse)
	return ctx.JSON(201, apiResponse)
}

func (menuController *MenuController) Update(ctx echo.Context) error {
	updateMenuRequest := request.UpdateMenuRequest{}

	err := ctx.Bind(&updateMenuRequest)
	if err != nil {
		fmt.Println("error binding")
		apiResponse := response.NewApiResponse("error", "failed create menu", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&updateMenuRequest)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed create menu", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	menuResponse, err := menuController.menuService.Update(updateMenuRequest)
	if err != nil {
		fmt.Println("error service")
		apiResponse := response.NewApiResponse("error", "failed create menu", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success update menu", menuResponse)
	return ctx.JSON(201, apiResponse)
}

func (menuController *MenuController) Get(ctx echo.Context) error {
	getMenuRequest := request.GetMenuRequest{}
	err := ctx.Bind(&getMenuRequest)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail menu", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&getMenuRequest)
	if err != nil {
		errorValidation := helper.FormatErrorValidation(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed get detail menu", errorValidation)
		return ctx.JSON(422, apiResponse)
	}

	menuResponse, err := menuController.menuService.Get(getMenuRequest)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail menu", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success get detail menu", menuResponse)
	return ctx.JSON(200, apiResponse)
}

func (menuController *MenuController) GetAll(ctx echo.Context) error {
	getAllMenuRequest := request.GetAllMenuRequest{}
	err := ctx.Bind(&getAllMenuRequest)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get menu", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	menuResponse, err := menuController.menuService.GetAll(getAllMenuRequest)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get menu", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success get menu", menuResponse)
	return ctx.JSON(200, apiResponse)
}
