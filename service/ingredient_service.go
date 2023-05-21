package service

import (
	"github.com/erp_app/models"
	"github.com/erp_app/repository"
	"github.com/erp_app/request"
	"github.com/erp_app/response"
)

type IngredientService interface {
	Create(createRequestIngredient request.CreateRequestIngredient) (response.IngredientResponse, error)
	Get(getDetailRequestIngredient request.GetDetailRequestIngredient) (response.IngredientResponse, error)
	GetAll(getAllRequestIngredient request.GetAllRequestIngredient) ([]response.IngredientResponse, error)
	Update(updateRequestIngredient request.UpdateRequestIngredient) (response.IngredientResponse, error)
	Delete(deleteRequestIngredient request.DeleteRequestIngredient) error
}

type ingredientService struct {
	ingredientRepository repository.IngredientRepository
}

func NewIngredientService(ingredientRepository repository.IngredientRepository) IngredientService {
	return &ingredientService{ingredientRepository: ingredientRepository}
}

func (ingredientService *ingredientService) Create(createRequestIngredient request.CreateRequestIngredient) (response.IngredientResponse, error) {
	res := response.IngredientResponse{}

	ingredient := models.Ingredient{}
	ingredient.Name = createRequestIngredient.Name

	ingredient, err := ingredientService.ingredientRepository.Create(ingredient)
	if err != nil {
		return res, err
	}

	res.Id = ingredient.Id
	res.Name = ingredient.Name

	return res, nil
}

func (ingredientService *ingredientService) Get(getDetailRequestIngredient request.GetDetailRequestIngredient) (response.IngredientResponse, error) {
	res := response.IngredientResponse{}
	ingredient, err := ingredientService.ingredientRepository.Find(getDetailRequestIngredient.Id)
	if err != nil {
		return res, err
	}

	res.Id = ingredient.Id
	res.Name = ingredient.Name

	return res, nil
}

func (ingredientService *ingredientService) GetAll(getAllRequestIngredient request.GetAllRequestIngredient) ([]response.IngredientResponse, error) {
	var listRes []response.IngredientResponse
	listIngredient, err := ingredientService.ingredientRepository.All(getAllRequestIngredient.Name)
	if err != nil {
		return listRes, err
	}

	if len(listIngredient) > 0 {
		for _, ingredient := range listIngredient {
			res := response.IngredientResponse{}
			res.Id = ingredient.Id
			res.Name = ingredient.Name

			listRes = append(listRes, res)
		}
	}

	return listRes, nil

}

func (ingredientService *ingredientService) Update(updateRequestIngredient request.UpdateRequestIngredient) (response.IngredientResponse, error) {
	res := response.IngredientResponse{}

	ingredient, err := ingredientService.ingredientRepository.Find(updateRequestIngredient.Id)
	if err != nil {
		return res, err
	}

	ingredient.Name = updateRequestIngredient.Name

	ingredient, err = ingredientService.ingredientRepository.Update(ingredient)
	if err != nil {
		return res, err
	}

	res.Id = ingredient.Id
	res.Name = ingredient.Name

	return res, nil
}

func (ingredientService *ingredientService) Delete(deleteRequestIngredient request.DeleteRequestIngredient) error {
	ingredient, err := ingredientService.ingredientRepository.Find(deleteRequestIngredient.Id)
	if err != nil {
		return err
	}

	err = ingredientService.ingredientRepository.Delete(ingredient)
	if err != nil {
		return err
	}

	return nil
}
