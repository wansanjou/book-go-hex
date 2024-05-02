package repository

import "gorm.io/gorm"

type bookRepositoryDB struct {
	db *gorm.DB
}

func NewBookRepositoryDB(db *gorm.DB) BookRepository  {
	return bookRepositoryDB{db: db}
}

func (b bookRepositoryDB) GetAll() ([]Book , error) {
	books := []Book{}
	err := b.db.Find(&books).Error
	if err != nil {
		return nil , err
	}

	return books , nil
}

func (b bookRepositoryDB) GetByID(id int) (*Book , error) {
	book := Book{}
	result := b.db.First(&book , id)
	if result.Error != nil {
		return nil , result.Error
	}

	return &book , nil
}

func (b bookRepositoryDB) CreateBook(book Book) (*Book , error)  {
	err := b.db.Create(&book).Error
	if err != nil {
		return nil, err
	}

	return &book , nil
}

func (b bookRepositoryDB) UpdateBook(id int , book Book) (*Book , error) {
	result := b.db.Where("id = ?", id).Updates(book)
	if result.Error != nil {
			return nil, result.Error
	}

	return &book, nil
}

func (b bookRepositoryDB) DeleteBook(id int) (*Book , error) {
	book := Book{}
	result := b.db.Model(&Book{}).Where("id = ?", id).Delete(&book)
	if result.Error != nil {
			return nil, result.Error
	}

	return nil, nil
}