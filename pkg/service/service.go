package service

import "pizza-world/pkg/repository"

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
