package services

import (
	"errors"
	"go-api-example/internal/models"
	"go-api-example/internal/repositories"
)

type UserService interface {
	CreateUser(req models.CreateUserRequest) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Age > 0 {
		user.Age = req.Age
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
