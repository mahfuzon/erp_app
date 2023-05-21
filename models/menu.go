package models

type Menu struct {
	Id          int
	Name        string
	CategoryId  int
	Category    Category
	Ingredients []MenuIngredient
}

func (menu *Menu) TableName() string {
	return "menus"
}
