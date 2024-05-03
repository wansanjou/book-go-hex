package service

import "gorm.io/gorm"

type PublisherResponse struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type PublisherService interface {
	GetPublisherAll() ([]PublisherResponse, error)
	GetPublisherByID(int) (*PublisherResponse, error)
	CreatePublisher(PublisherResponse) (*PublisherResponse, error)
	UpdatePublisher(int, PublisherResponse) (*PublisherResponse, error)
	DeletePublisher(int) (*PublisherResponse, error)
}