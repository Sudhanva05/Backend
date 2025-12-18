package service

import (
	"context"
	"time"

	db "github.com/Sudhanva05/Backend/db/sqlc"
	"github.com/Sudhanva05/Backend/internal/models"
)

// UserService handles user business logic
type UserService struct {
	queries *db.Queries
}

// NewUserService creates a new UserService
func NewUserService(queries *db.Queries) *UserService {
	return &UserService{
		queries: queries,
	}
}

// CreateUser creates a new user in PostgreSQL
func (s *UserService) CreateUser(req models.CreateUserRequest) (models.UserResponse, error) {

	// 1. Parse DOB string into time.Time
	dob, err := req.DOBTime()
	if err != nil {
		return models.UserResponse{}, err
	}

	// 2. Insert user using SQLC
	user, err := s.queries.CreateUser(
		context.Background(),
		db.CreateUserParams{
			Name: req.Name,
			Dob:  dob,
		},
	)
	if err != nil {
		return models.UserResponse{}, err
	}

	// 3. Build API response
	return models.UserResponse{
		ID:   int64(user.ID),
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(id int64) (models.UserResponse, error) {

	user, err := s.queries.GetUserByID(
		context.Background(),
		int32(id),
	)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   int64(user.ID),
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

// GetAllUsers fetches all users
func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {

	users, err := s.queries.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}

	var result []models.UserResponse
	for _, user := range users {
		result = append(result, models.UserResponse{
			ID:   int64(user.ID),
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  calculateAge(user.Dob),
		})
	}

	return result, nil
}

// calculateAge calculates age dynamically
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
