package repository

import "github.com/jmoiron/sqlx"

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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
