package service

import (
	"errors"
	"wansanjou/logs"
	"wansanjou/repository"
)

type publisherService struct {
	publisher_repo repository.PublisherRepository
}

func NewPublisherService(publisher_repo repository.PublisherRepository) PublisherService {
	return publisherService{publisher_repo: publisher_repo}
}

func (ps publisherService) GetPublisherAll() ([]PublisherResponse, error) {
	publishers , err := ps.publisher_repo.GetAll()
	if err != nil {
		return nil , err
	}

	publishers_responses := []PublisherResponse{}
	for _ , publisher := range publishers {
		publishers_response := PublisherResponse{
			Name: publisher.Name,
			Address: publisher.Address,
			Phone: publisher.Phone,
			Email: publisher.Email,
		}
		publishers_responses = append(publishers_responses , publishers_response)
	} 

	return publishers_responses , nil
}

func (ps publisherService) GetPublisherByID(id int) (*PublisherResponse, error)  {
	publisher , err := ps.publisher_repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	publishers_response := PublisherResponse{
		Name: publisher.Name,
		Address: publisher.Address,
		Phone: publisher.Phone,
		Email: publisher.Email,
	}

	return &publishers_response , nil
}

func (ps publisherService) CreatePublisher(publishers_res PublisherResponse) (*PublisherResponse, error)  {
	if publishers_res.Name == ""  {
		logs.Error("Pls enter publisher name")
	}

	if publishers_res.Address == ""  {
		logs.Error("Pls enter publisher address")
	}

	if publishers_res.Phone == ""  {
		logs.Error("Pls enter publisher phone")
	}

	if publishers_res.Email == ""  {
		logs.Error("Pls enter publisher email")
	}

	publisher := repository.Publisher{
		Name: publishers_res.Name,
		Address: publishers_res.Address,
		Phone: publishers_res.Phone,
		Email: publishers_res.Email,
	}

	i_publisher , err := ps.publisher_repo.CreatePublisher(publisher)
	if err != nil {
		logs.Error(err)
	}

	response := PublisherResponse{
		Name: i_publisher.Name,
		Address: i_publisher.Address,
		Phone: i_publisher.Phone,
		Email: i_publisher.Email,
	}

	return &response , nil
}

func (ps publisherService) UpdatePublisher(id int, publishers_res PublisherResponse) (*PublisherResponse, error) {
	if id == 0 {
			logs.Error("Invalid ID")
			return nil, errors.New("Invalid ID")
	}


	publisher := repository.Publisher {
			Name:        publishers_res.Name,
			Address: 		 publishers_res.Address,
			Phone:       publishers_res.Phone,
			Email:       publishers_res.Email,
	}

	u_publisher, err := ps.publisher_repo.UpdatePublisher(id, publisher)
	if err != nil {
			logs.Error(err)
			return nil, errors.New("Failed to update publisher")
	}

	response := PublisherResponse{
			Name:        u_publisher.Name,
			Address: u_publisher.Address,
			Phone:       u_publisher.Phone,
			Email:       u_publisher.Email,
	}

	return &response, nil
}



func (ps publisherService) DeletePublisher(id int) (*PublisherResponse, error)  {
	if id == 0 {
		logs.Error("Invalid id")
	}

	_ , err := ps.publisher_repo.DeletePublisher(id)
	if err != nil {
		return nil, err
	}

	return nil , nil
}