package recipe

type RecipeList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id       int
	UserId   int
	RecipeId int
}

type Ingredient struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Meat        string `json:"meat"`
	Seafood     string `json:"seafood"`
}

type IngredientsList struct {
	Id           int
	RecipeId     int
	IngredientId int
}
