package service

import (
	"capstone/internal/model"
	"capstone/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user model.User) error {
	return s.repo.CreateUser(user)
}
