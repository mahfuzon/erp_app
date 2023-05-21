package response

type MenuResponse struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	CategoryId  int              `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Ingredients []RecipeResponse `json:"ingredients"`
}

type RecipeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Qty  string `json:"qty"`
}
