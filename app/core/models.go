package core

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
