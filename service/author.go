package service

import "gorm.io/gorm"

type AuthorResponse struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthorService interface {
	GetAuthorAll() ([]AuthorResponse, error)
	GetAuthorByID(int) (*AuthorResponse, error)
	CreateAuthor(AuthorResponse) (*AuthorResponse, error)
	UpdateAuthor(int, AuthorResponse) (*AuthorResponse, error)
	DeleteAuthor(int) (*AuthorResponse, error)
}

