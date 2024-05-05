package service

import "gorm.io/gorm"

// type UserRequest struct {
// 	Email string `json:"email"`
// 	Firstname string `json:"firstname"`
// 	Lastname string `json:"lastname"`
// }

type UserResponse struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type UserService interface {
	GetUserAll() ([]UserResponse, error)
	GetUserByID(int) (*UserResponse, error)
	CreateUser(UserResponse) (*UserResponse, error)
	UpdateUser(int, UserResponse) (*UserResponse, error)
	DeleteUser(int) (*UserResponse, error)
	LoginUser(UserResponse)  (*UserResponse , error)
}