package service

import (
	"to-do-list/internal/domain/models"
	"to-do-list/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) Login(user *models.Login) (*models.Login, error) {
	return s.repo.Login(user)
}
