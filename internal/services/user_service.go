package services

import (
	"blog/internal/models"
	"blog/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) CreateUser(user *models.User) (int64, error) {
	return s.userRepository.CreateUser(user)
}
