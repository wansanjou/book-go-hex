package repository

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `db:"name"`
	Email string `db:"email"`
}

type AuthorRepository interface {
	GetAll() ([]Author , error)
	GetByID(int) (*Author , error)
	CreateAuthor(Author) (*Author , error)
	UpdateAuthor(int , Author) (*Author , error)
	DeleteAuthor(int) (*Author , error)
}