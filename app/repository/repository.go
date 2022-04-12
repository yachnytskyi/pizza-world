package repository

import (
	"database/sql"

	"github.com/yachnytskyi/pizza-world/app/core"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
}

type Recipe interface {
}

type Ingredient interface {
}

type Repository struct {
	Authorization
	Recipe
	Ingredient
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
