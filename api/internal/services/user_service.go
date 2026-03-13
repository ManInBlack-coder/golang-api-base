package services

import (
	"fmt"
	"sync"
	"time"

	"golang-api/internal/models"
)

// UserService handles user-related business logic
type UserService struct {
	users  map[string]models.User
	mu     sync.RWMutex
	nextID int
}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	service := &UserService{
		users:  make(map[string]models.User),
		nextID: 1,
	}

	// Add some sample data
	service.seedData()

	return service
}

// seedData adds initial sample users
func (s *UserService) seedData() {
	sampleUsers := []models.User{
		{
			ID:        "1",
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		{
			ID:        "2",
			Name:      "Jane Smith",
			Email:     "jane.smith@example.com",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}

	for _, user := range sampleUsers {
		s.users[user.ID] = user
	}
	s.nextID = 3
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// GetUserByID returns a user by ID
func (s *UserService) GetUserByID(id string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req models.CreateUserRequest) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Format(time.RFC3339)
	user := models.User{
		ID:        fmt.Sprintf("%d", s.nextID),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.users[user.ID] = user
	s.nextID++

	return user
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id string, req models.UpdateUserRequest) (models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, false
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	s.users[id] = user
	return user, true
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[id]
	if !exists {
		return false
	}

	delete(s.users, id)
	return true
}
