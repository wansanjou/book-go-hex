package repository

import "gorm.io/gorm"

type publisherRepositoryDB struct {
	db *gorm.DB
}

func NewPublisherRepositoryDB(db *gorm.DB) PublisherRepository  {
	return publisherRepositoryDB{db: db}
}

func (p publisherRepositoryDB) GetAll() ([]Publisher , error)  {
	publishers := []Publisher{}
	err := p.db.Find(&publishers).Error
	if err != nil {
		return nil , err
	}

	return publishers , nil
}

func (p publisherRepositoryDB) GetByID(id int) (*Publisher , error) {
	publisher := Publisher{}
	result := p.db.First(&publisher , id)
	if result.Error != nil {
		return nil , result.Error
	}

	return &publisher , nil
}

func (p publisherRepositoryDB) CreatePublisher(publisher Publisher) (*Publisher , error)  {
	err := p.db.Create(&publisher).Error
	if err != nil {
		return nil, err
	}

	return &publisher , nil
}

func (p publisherRepositoryDB) UpdatePublisher(id int , publisher Publisher) (*Publisher , error) {
	result := p.db.Where("id = ?", id).Updates(publisher)
	if result.Error != nil {
			return nil, result.Error
	}

	return &publisher, nil
}

func (p publisherRepositoryDB) DeletePublisher(id int) (*Publisher , error) {
	publisher := Publisher{}
	result := p.db.Model(&Publisher{}).Where("id = ?", id).Delete(&publisher)
	if result.Error != nil {
			return nil, result.Error
	}

	return nil, nil
}