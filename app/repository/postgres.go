package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable        = "users"
	recipesTable      = "recipes"
	ingredientsTable  = "ingredients"
	recipeUsers       = "recipes_users"
	recipeIngredients = "recipe_ingredients"
)

type Config struct {
	DB_HOST  string
	DB_PORT  string
	DB_NAME  string
	DB_USER  string
	DB_PASS  string
	SSL_MODE string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("db_host=%s db_port=%s db_name=%s db_user=%s db_pass=%s sslmode=%s",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME, cfg.DB_USER, cfg.DB_PASS, cfg.SSL_MODE))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
