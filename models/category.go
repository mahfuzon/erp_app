package models

type Category struct {
	Id   int
	Name string
}

func (category *Category) TableName() string {
	return "categories"
}
