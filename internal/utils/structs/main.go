package structs

type Recipe struct {
	Title             string `json:"title"`
	Ingredients       string `json:"ingredients"`
	Instructions      string `json:"instructions"`
	DietaryCompliance string `json:"dietaryCompliance"`
}

type RecipesResponse struct {
	Recipes []Recipe `json:"recipes"`
}
