package request

type CreateRequestCategory struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRequestCategory struct {
	Name string `json:"name" validate:"required"`
	Id   int    `param:"id" validate:"required"`
}

type GetDetailRequestCategory struct {
	Id int `param:"id" validate:"required"`
}

type GetAllRequestCategory struct {
	Name string `query:"name"`
}

type DeleteRequestCategory struct {
	Id int `param:"id" validate:"required"`
}
