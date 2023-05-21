package service

import (
	"fmt"
	"github.com/erp_app/models"
	"github.com/erp_app/repository"
	"github.com/erp_app/request"
	"github.com/erp_app/response"
)

type MenuService interface {
	Create(createMenuRequest request.CreateMenuRequest) (response.MenuResponse, error)
	Update(updateMenuRequest request.UpdateMenuRequest) (response.MenuResponse, error)
	Get(getMenuRequest request.GetMenuRequest) (response.MenuResponse, error)
	GetAll(getAllMenuRequest request.GetAllMenuRequest) ([]response.MenuResponse, error)
	Delete(deleteRequestIngredient request.DeleteMenuRequest) error
}

type menuService struct {
	menuRepository     repository.MenuRepository
	categoryRepository repository.CategoryRepository
}

func NewMenuService(menuRepository repository.MenuRepository, categoryRepository repository.CategoryRepository) MenuService {
	return &menuService{
		menuRepository:     menuRepository,
		categoryRepository: categoryRepository,
	}
}

func (menuService *menuService) Delete(deleteMenuRequest request.DeleteMenuRequest) error {
	menu, err := menuService.menuRepository.Find(deleteMenuRequest.Id)
	if err != nil {
		return err
	}

	err = menuService.menuRepository.Delete(menu)
	if err != nil {
		return err
	}

	return nil
}

func (menuService *menuService) Create(createMenuRequest request.CreateMenuRequest) (response.MenuResponse, error) {

	res := response.MenuResponse{}

	category, err := menuService.categoryRepository.Find(createMenuRequest.CategoryId)
	if err != nil {
		return res, err
	}

	menu := models.Menu{}
	menu.Name = createMenuRequest.Name
	menu.CategoryId = createMenuRequest.CategoryId
	menu, err = menuService.menuRepository.Create(menu)
	if err != nil {
		return res, err
	}

	categoryResponse := response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	res.Id = menu.Id
	res.Name = menu.Name
	res.CategoryId = menu.CategoryId
	res.Category = categoryResponse

	return res, nil
}

func (menuService *menuService) Update(updateMenuRequest request.UpdateMenuRequest) (response.MenuResponse, error) {
	res := response.MenuResponse{}

	menu, err := menuService.menuRepository.Find(updateMenuRequest.Id)
	if err != nil {
		return res, err
	}

	category, err := menuService.categoryRepository.Find(updateMenuRequest.CategoryId)
	if err != nil {
		return res, err
	}

	menu.Name = updateMenuRequest.Name
	menu.CategoryId = updateMenuRequest.CategoryId

	menu, err = menuService.menuRepository.Update(menu)
	if err != nil {
		return res, err
	}

	categoryResponse := response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	res.Id = menu.Id
	res.Name = menu.Name
	res.CategoryId = menu.CategoryId
	res.Category = categoryResponse

	return res, nil
}

func (menuService *menuService) Get(getMenuRequest request.GetMenuRequest) (response.MenuResponse, error) {
	res := response.MenuResponse{}
	menu, err := menuService.menuRepository.Find(getMenuRequest.Id)
	if err != nil {
		return res, err
	}

	categoryRes := response.CategoryResponse{
		Id:   menu.Category.Id,
		Name: menu.Category.Name,
	}

	res.Id = menu.Id
	res.Name = menu.Name
	res.CategoryId = menu.CategoryId
	res.Category = categoryRes

	return res, nil
}

func (menuService *menuService) GetAll(getAllMenuRequest request.GetAllMenuRequest) ([]response.MenuResponse, error) {
	var listMenuResponse []response.MenuResponse

	listMenu, err := menuService.menuRepository.All(getAllMenuRequest.Name)
	if err != nil {
		return listMenuResponse, err
	}

	fmt.Println(listMenu)

	if len(listMenu) > 0 {
		for _, menu := range listMenu {

			var listRecipeResponse []response.RecipeResponse
			if len(menu.Ingredients) > 0 {
				for _, ingredient := range menu.Ingredients {
					recipeResponse := response.RecipeResponse{
						Id:   ingredient.IngredientId,
						Name: ingredient.Ingredient.Name,
						Qty:  ingredient.Qty,
					}

					listRecipeResponse = append(listRecipeResponse, recipeResponse)

				}
			}

			categoryRes := response.CategoryResponse{
				Id:   menu.Category.Id,
				Name: menu.Category.Name,
			}

			res := response.MenuResponse{}
			res.Id = menu.Id
			res.Name = menu.Name
			res.CategoryId = menu.CategoryId
			res.Category = categoryRes
			res.Ingredients = listRecipeResponse

			listMenuResponse = append(listMenuResponse, res)
		}
	}

	//fmt.Println(listMenuResponse)
	return listMenuResponse, nil
}
