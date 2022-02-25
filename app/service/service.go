package service

import "github.com/yachnytskyi/pizza-world/app/repository"

type Authorization interface {
}

type RecipeList interface {
}

type IngredientList interface {
}

type Service struct {
	Authorization
	RecipeList
	IngredientList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
