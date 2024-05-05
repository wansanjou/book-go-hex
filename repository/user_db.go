package repository

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryFB(db *gorm.DB) UserRepository  {
	return userRepositoryDB{db: db}
}

func (u userRepositoryDB) GetAll() ([]User , error) {
	users := []User{}
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users , nil
}

func (u userRepositoryDB) GetByID(id int) (*User , error)  {
	user := User{}
	err := u.db.First(&user , id).Error
	if err != nil {
		return nil, err
	}

	return &user , nil
}

func (u userRepositoryDB) CreateUser(user User) (*User , error)  {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return &user , nil
}

func (u userRepositoryDB) UpdateUser(id int , user User) (*User , error)  {
	result := u.db.Where("id = ?", id).Updates(user)
	if result.Error != nil {
			return nil, result.Error
	}

	return &user, nil
}

func (u userRepositoryDB) DeleteUser(id int) (*User , error) {
	user := User{}
	result := u.db.Model(&User{}).Where("id = ?", id).Delete(&user)
	if result.Error != nil {
			return nil, result.Error
	}

	return nil, nil
}

func (u userRepositoryDB) LoginUser(user User) (*User, error) {
	selectedUser := new(User)
	result := u.db.Where("email = ?", user.Email).First(selectedUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return selectedUser, nil
}