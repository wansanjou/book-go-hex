package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
	Password string
	Firstname string
	Lastname string
}

type UserRepository interface {
	GetAll() ([]User , error)
	GetByID(int) (*User , error)
	CreateUser(User) (*User , error)
	UpdateUser(int , User) (*User , error)
	DeleteUser(int) (*User , error)
	LoginUser(User)  (*User , error)
}