package service

type BookResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	Genres      string  `json:"genres"`
	AuthorID    uint    `json:"author_id"`
	PublisherID uint    `json:"publisher_id"`
}

type BookService interface {
	GetBookAll() ([]BookResponse, error)
	GetBookByID(int) (*BookResponse, error)
	CreateBook(BookResponse) (*BookResponse, error)
	UpdateBook(int, BookResponse) (*BookResponse, error)
	DeleteBook(int) (*BookResponse, error)
}
