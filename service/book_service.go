package service

import (
	"errors"
	"wansanjou/logs"
	"wansanjou/repository"
)

type bookService struct {
	book_repo repository.BookRepository
}

func NewBookService(book_repo repository.BookRepository) BookService {
	return bookService{book_repo: book_repo}
}

func (bs bookService) GetBookAll() ([]BookResponse, error) {
	books , err := bs.book_repo.GetAll()
	if err != nil {
		return nil , err
	}

	book_responses := []BookResponse{}
	for _ , book := range books {
		book_response := BookResponse{
			ID : book.ID,
			Name: book.Name,
			Description: book.Description,
			Price: book.Price,
			Stock: book.Stock,
			Genres: book.Genres,
			AuthorID: book.AuthorID,
			PublisherID: book.PublisherID,
		}
		book_responses = append(book_responses , book_response)
	} 

	return book_responses , nil
}

func (bs bookService) GetBookByID(id int) (*BookResponse, error)  {
	book , err := bs.book_repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	book_response := BookResponse{
		ID : book.ID,
		Name: book.Name,
		Description: book.Description,
		Price: book.Price,
		Stock: book.Stock,
		Genres: book.Genres,
		AuthorID: book.AuthorID,
		PublisherID: book.PublisherID,
	}

	return &book_response , nil
}

func (bs bookService) CreateBook(book_res BookResponse) (*BookResponse, error)  {
	if book_res.Stock <= 0  {
		logs.Error("Stock must least than 1")
	}

	if book_res.Price <= 0  {
		logs.Error("Price must more than 0")
	}

	book := repository.Book{
		Name: book_res.Name,
		Description: book_res.Description,
		Price: book_res.Price,
		Stock: book_res.Stock,
		Genres: book_res.Genres,
		AuthorID: book_res.AuthorID,
		PublisherID: book_res.PublisherID,
	}

	i_book , err := bs.book_repo.CreateBook(book)
	if err != nil {
		logs.Error(err)
	}

	response := BookResponse{
		Name: i_book.Name,
		Description: i_book.Description,
		Price: i_book.Price,
		Stock: i_book.Stock,
		Genres: i_book.Genres,
		AuthorID: i_book.AuthorID,
		PublisherID: i_book.PublisherID,
	}

	return &response , nil
}

func (bs bookService) UpdateBook(id int, book_res BookResponse) (*BookResponse, error) {
	if id == 0 {
			logs.Error("Invalid ID")
			return nil, errors.New("Invalid ID")
	}

	if book_res.Stock <= 0 {
			logs.Error("Stock must be at least 1")
			return nil, errors.New("Stock must be at least 1")
	}

	if book_res.Price <= 0 {
			logs.Error("Price must be more than 0")
			return nil, errors.New("Price must be more than 0")
	}

	book := repository.Book{
			Name:        book_res.Name,
			Description: book_res.Description,
			Price:       book_res.Price,
			Stock:       book_res.Stock,
			Genres:      book_res.Genres,
	}

	if book_res.AuthorID != 0 {
			book.AuthorID = book_res.AuthorID
	}
	if book_res.PublisherID != 0 {
			book.PublisherID = book_res.PublisherID
	}

	u_book, err := bs.book_repo.UpdateBook(id, book)
	if err != nil {
			logs.Error(err)
			return nil, errors.New("Failed to update book")
	}

	response := BookResponse{
			Name:        u_book.Name,
			Description: u_book.Description,
			Price:       u_book.Price,
			Stock:       u_book.Stock,
			Genres:      u_book.Genres,
			AuthorID:    u_book.AuthorID,
			PublisherID: u_book.PublisherID,
	}

	return &response, nil
}



func (bs bookService) DeleteBook(id int) (*BookResponse, error)  {
	if id == 0 {
		logs.Error("Invalid id")
	}

	_ , err := bs.book_repo.DeleteBook(id)
	if err != nil {
		return nil, err
	}

	return nil , nil
}