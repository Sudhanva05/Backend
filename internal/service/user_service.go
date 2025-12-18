package service

import (
	"errors"
	"time"

	"github.com/Sudhanva05/Backend/internal/models"
)

// UserService contains business logic related to users
type UserService struct {
	users  map[int64]models.UserResponse
	nextID int64
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int64]models.UserResponse),
		nextID: 1,
	}
}

// CreateUser handles business logic for creating a user
func (s *UserService) CreateUser(req models.CreateUserRequest) models.UserResponse {

	dob, _ := time.Parse("2006-01-02", req.DOB)
	age := calculateAge(dob)

	user := models.UserResponse{
		ID:   s.nextID,
		Name: req.Name,
		DOB:  req.DOB,
		Age:  age,
	}

	s.users[s.nextID] = user
	s.nextID++

	return user
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(id int64) (models.UserResponse, error) {
	user, exists := s.users[id]
	if !exists {
		return models.UserResponse{}, errors.New("user not found")
	}
	return user, nil
}

// calculateAge calculates age from date of birth
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() []models.UserResponse {

	users := make([]models.UserResponse, 0, len(s.users))

	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}
