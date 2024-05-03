package service

import (
	"wansanjou/logs"
	"wansanjou/repository"
)

type authorService struct {
	author_repo repository.AuthorRepository
}

func  NewAuthorService(author_repo repository.AuthorRepository) AuthorService {
	return authorService{author_repo: author_repo}
}

func (as authorService) GetAuthorAll() ([]AuthorResponse, error) {
	authors , err := as.author_repo.GetAll()
	if err != nil {
		return nil, err
	}

	author_responses := []AuthorResponse{}
	for _ , author := range authors {
		author_response := AuthorResponse{
			Name: author.Name,
			Email: author.Email,
		}
		author_responses =  append(author_responses , author_response)
	}

	return author_responses , nil
}

func (as authorService) GetAuthorByID(id int) (*AuthorResponse, error) {
	author , err := as.author_repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	author_response := AuthorResponse{
		Name: author.Name,
		Email: author.Email,
	}


	return &author_response , nil
}

func (as authorService) CreateAuthor(author_res AuthorResponse) (*AuthorResponse, error) {
	if author_res.Name == "" {
		logs.Error("Pls enter Author name")
	}

	if author_res.Email == "" {
		logs.Error("Pls enter Author email")
	}

	author := repository.Author{
		Name: author_res.Name,
		Email: author_res.Email,
	}

	i_author , err := as.author_repo.CreateAuthor(author)
	if err != nil {
		logs.Error("Insert Error!")
	}

	response := AuthorResponse{
		Name: i_author.Name,
		Email: i_author.Email,
	}

	return &response , nil
}

func (as authorService) UpdateAuthor(id int, author_res AuthorResponse) (*AuthorResponse, error) {
	if id == 0 {
		logs.Error("invalid id")
	}

	author := repository.Author{
		Name: author_res.Name,
		Email: author_res.Email,
	}

	u_author , err := as.author_repo.UpdateAuthor(id , author)
	if err != nil {
		logs.Error("Insert error")
	}

	response := AuthorResponse{
		Name: u_author.Name,
		Email: u_author.Email,

	}

	return &response , nil
}

func (as authorService) DeleteAuthor(id int) (*AuthorResponse, error) {
	if id == 0 {
		logs.Error("invalid id")
	}

	_ , err := as.author_repo.DeleteAuthor(id)
	if err != nil {
		return nil, err
	}

	return nil , nil
}