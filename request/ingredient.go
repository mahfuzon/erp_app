package request

type CreateRequestIngredient struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRequestIngredient struct {
	Name string `json:"name" validate:"required"`
	Id   int    `param:"id" validate:"required"`
}

type GetDetailRequestIngredient struct {
	Id int `param:"id" validate:"required"`
}

type GetAllRequestIngredient struct {
	Name string `query:"name"`
}

type DeleteRequestIngredient struct {
	Id int `param:"id" validate:"required"`
}
