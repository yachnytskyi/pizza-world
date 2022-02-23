package repository

type Authorization interface {
}

type RecipeList interface {
}

type IngredientList interface {
}

type Repository struct {
	Authorization
	RecipeList
	IngredientList
}

func NewRepository() *Repository {
	return &Repository{}
}
