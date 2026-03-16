package repository

import (
	"golang-api/internal/models"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id int, user models.User) (models.User, error)
	Delete(id int) error
}
