package core

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Recipe struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RecipeUsers struct {
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

type RecipeIngredients struct {
	Id           int
	RecipeId     int
	IngredientId int
}
