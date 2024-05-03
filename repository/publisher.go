package repository

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name    string `db:"name"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
	Email   string `db:"email"`
}

type PublisherRepository interface {
	GetAll() ([]Publisher , error)
	GetByID(int) (*Publisher , error)
	CreatePublisher(Publisher) (*Publisher , error)
	UpdatePublisher(int , Publisher) (*Publisher , error)
	DeletePublisher(int) (*Publisher , error)
}