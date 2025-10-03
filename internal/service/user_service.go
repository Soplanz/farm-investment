package service

import (
	"farm-investment/internal/dto"
	"farm-investment/internal/models"
	"farm-investment/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(req dto.CreateUserRequest) (models.User, error)
}

type userService struct {
	repo repository.UserRepository // ini interface dari repository layer
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (models.User, error) {
	user := models.User{
		Username: req.Name,
		Email:    req.Email,
	}
	return s.repo.Create(user)
}
