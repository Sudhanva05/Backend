package service

import (
	"time"

	"github.com/Sudhanva05/Backend/internal/models"
)

// UserService contains business logic related to users
type UserService struct {
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser handles business logic for creating a user
func (s *UserService) CreateUser(req models.CreateUserRequest) models.UserResponse {

	// Parse DOB string to time.Time
	dob, _ := time.Parse("2006-01-02", req.DOB)

	age := calculateAge(dob)

	return models.UserResponse{
		ID:   1, // temporary; DB will generate later
		Name: req.Name,
		DOB:  req.DOB,
		Age:  age,
	}
}

// calculateAge calculates age from date of birth
func calculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// If birthday hasn't occurred yet this year, subtract 1
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
