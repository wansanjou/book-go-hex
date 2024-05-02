package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float32 `db:"price"`
	Stock       int 		`db:"stock"`
	Genres      string 	`db:"genres"`
	AuthorID 		uint 		`db:"author_id"`
	Author      Author 
	PublisherID uint 		`db:"publisher_id"`
	Publisher   Publisher
}

type BookRepository interface {
	GetAll() ([]Book , error)
	GetByID(int) (*Book , error)
	CreateBook(Book) (*Book , error)
	UpdateBook(int , Book) (*Book , error)
	DeleteBook(int) (*Book , error)
}