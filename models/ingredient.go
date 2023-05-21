package models

type Ingredient struct {
	Id   int
	Name string
}

func (ingredient *Ingredient) TableName() string {
	return "ingredients"
}
