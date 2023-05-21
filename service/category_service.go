package service

import (
	"github.com/erp_app/models"
	"github.com/erp_app/repository"
	"github.com/erp_app/request"
	"github.com/erp_app/response"
)

type CategoryService interface {
	Create(createRequestCategory request.CreateRequestCategory) (response.CategoryResponse, error)
	Get(getDetailCategoryRequest request.GetDetailRequestCategory) (response.CategoryResponse, error)
	GetAll(getAllCategoryRequest request.GetAllRequestCategory) ([]response.CategoryResponse, error)
	Update(updateRequestCategory request.UpdateRequestCategory) (response.CategoryResponse, error)
	Delete(deleteRequestCategory request.DeleteRequestCategory) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (categoryService *categoryService) Create(createRequestCategory request.CreateRequestCategory) (response.CategoryResponse, error) {
	res := response.CategoryResponse{}

	category := models.Category{}
	category.Name = createRequestCategory.Name

	category, err := categoryService.categoryRepository.Create(category)
	if err != nil {
		return res, err
	}

	res.Id = category.Id
	res.Name = category.Name

	return res, nil
}

func (categoryService *categoryService) Get(getDetailCategoryRequest request.GetDetailRequestCategory) (response.CategoryResponse, error) {
	res := response.CategoryResponse{}
	category, err := categoryService.categoryRepository.Find(getDetailCategoryRequest.Id)
	if err != nil {
		return res, err
	}

	res.Id = category.Id
	res.Name = category.Name

	return res, nil
}

func (categoryService *categoryService) GetAll(getAllCategoryRequest request.GetAllRequestCategory) ([]response.CategoryResponse, error) {
	var listRes []response.CategoryResponse
	listCategory, err := categoryService.categoryRepository.All(getAllCategoryRequest.Name)
	if err != nil {
		return listRes, err
	}

	if len(listCategory) > 0 {
		for _, category := range listCategory {
			res := response.CategoryResponse{}
			res.Id = category.Id
			res.Name = category.Name

			listRes = append(listRes, res)
		}
	}

	return listRes, nil

}

func (categoryService *categoryService) Update(updateRequestCategory request.UpdateRequestCategory) (response.CategoryResponse, error) {
	res := response.CategoryResponse{}

	category, err := categoryService.categoryRepository.Find(updateRequestCategory.Id)
	if err != nil {
		return res, err
	}

	category.Name = updateRequestCategory.Name

	category, err = categoryService.categoryRepository.Update(category)
	if err != nil {
		return res, err
	}

	res.Id = category.Id
	res.Name = category.Name

	return res, nil
}

func (categoryService *categoryService) Delete(deleteRequestCategory request.DeleteRequestCategory) error {
	category, err := categoryService.categoryRepository.Find(deleteRequestCategory.Id)
	if err != nil {
		return err
	}

	err = categoryService.categoryRepository.Delete(category)
	if err != nil {
		return err
	}

	return nil
}
