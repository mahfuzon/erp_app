package models

type MenuIngredient struct {
	Id           int
	MenuId       int
	IngredientId int
	Qty          string
	Ingredient   Ingredient
}

func (recipe *MenuIngredient) TableName() string {
	return "recipes"
}
