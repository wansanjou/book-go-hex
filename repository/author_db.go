package repository

import "gorm.io/gorm"

type authorRepositoryDB struct {
	db *gorm.DB
}

func NewAuthorRepostioryDB(db *gorm.DB) AuthorRepository {
	return authorRepositoryDB{db: db}
}

func (a authorRepositoryDB) GetAll() ([]Author , error)  {
	authors := []Author{}
	err := a.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}

	return authors , nil
}

func (a authorRepositoryDB) GetByID(id int) (*Author , error)  {
	author := Author{}
	err := a.db.First(&author , id).Error
	if err != nil {
		return nil, err
	}

	return &author , nil
}

func (a authorRepositoryDB) CreateAuthor(author Author) (*Author , error)  {
	err := a.db.Create(author).Error
	if err != nil {
		return nil, err
	}

	return &author , nil
}

func (a authorRepositoryDB) UpdateAuthor(id int , author Author) (*Author , error)  {
	result := a.db.Where("id = ?", id).Updates(author)
	if result.Error != nil {
			return nil, result.Error
	}

	return &author, nil
}

func (a authorRepositoryDB) DeleteAuthor(id int) (*Author , error)  {
	author := Author{}
	result := a.db.Model(&Author{}).Where("id = ?", id).Delete(&author)
	if result.Error != nil {
			return nil, result.Error
	}

	return nil, nil
}