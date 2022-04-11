package service

import (
	"github.com/yachnytskyi/pizza-world/app/core"
	"github.com/yachnytskyi/pizza-world/app/repository"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
}

type Recipe interface {
}

type Ingredient interface {
}

type Service struct {
	Authorization
	Recipe
	Ingredient
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
