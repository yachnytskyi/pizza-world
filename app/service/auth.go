package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/yachnytskyi/pizza-world/app/core"
	"github.com/yachnytskyi/pizza-world/app/repository"
)

const salt = "dashuuvcxn4181-528124!@($&!)#_"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user core.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
