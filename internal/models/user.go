package models

type CreateUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}
