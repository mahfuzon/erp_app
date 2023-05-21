package request

type CreateRecipeRequest struct {
	IngredientId int    `json:"ingredient_id" validate:"required,gte=1"`
	MenuId       int    `param:"menu_id" validate:"required,gte=1"`
	Qty          string `json:"qty" validate:"required"`
}

type UpdateRecipeRequest struct {
	Id           int    `param:"id" validate:"required"`
	IngredientId int    `json:"ingredient_id" validate:"required,gte=1"`
	MenuId       int    `param:"menu_id" validate:"required,gte=1"`
	Qty          string `json:"qty" validate:"required"`
}

type DeleteRecipeRequest struct {
	Id     int `param:"id" validate:"required"`
	MenuId int `param:"id" validate:"required"`
}
