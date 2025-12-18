package models

import "time"

// CreateUserRequest represents POST /users payload
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// DOBTime converts DOB string to time.Time
func (r CreateUserRequest) DOBTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.DOB)
}

// UserResponse represents API response
type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
