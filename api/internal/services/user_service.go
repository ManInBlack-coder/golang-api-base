package services

import (
	"golang-api/internal/models"
	"golang-api/internal/repository"
)

// UserService handles user-related business logic
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

// GetUserByID returns a user by ID
func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req models.CreateUserRequest) (models.User, error) {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	return s.repo.Create(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id int, req models.UpdateUserRequest) (models.User, error) {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	return s.repo.Update(id, user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
