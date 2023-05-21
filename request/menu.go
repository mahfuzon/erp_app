package request

type CreateMenuRequest struct {
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required,gte=1"`
}

type UpdateMenuRequest struct {
	Id         int    `param:"id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required,gte=1"`
}

type GetMenuRequest struct {
	Id int `param:"id" validate:"required"`
}

type GetAllMenuRequest struct {
	Name string `query:"name"`
}

type DeleteMenuRequest struct {
	Id int `param:"id" validate:"required"`
}
