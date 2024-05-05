package service

import (
	"errors"
	"wansanjou/logs"
	"wansanjou/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	user_repo repository.UserRepository
}

func NewUserService(user_repo repository.UserRepository) UserService {
	return userService{user_repo: user_repo}
}

func (us userService) GetUserAll() ([]UserResponse, error)  {
	users , err := us.user_repo.GetAll()
	if err != nil {
		return nil , err
	}

	users_responses := []UserResponse{}
	for _ , user := range users {
		users_response := UserResponse{
			Email: user.Email,
			Firstname: user.Firstname,
			Lastname: user.Lastname,
		}
		users_responses = append(users_responses , users_response)
	} 

	return users_responses , nil
}

func (us userService) GetUserByID(id int) (*UserResponse, error)  {
	user , err := us.user_repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	users_responses := UserResponse{
		Email: user.Email,
		Firstname: user.Firstname,
		Lastname: user.Lastname,
	}

	return &users_responses , nil
}

func (us userService) CreateUser(user_res UserResponse) (*UserResponse, error)  {
	if user_res.Email == ""  {
		return nil, errors.New("Please enter user Email")
	}

	if user_res.Password == ""  {
		return nil, errors.New("Please enter user Password")
	}

	if user_res.Firstname == ""  {
		return nil, errors.New("Please enter user Firstname")
	}

	if user_res.Lastname == ""  {
		return nil, errors.New("Please enter user Lastname")
	}

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(user_res.Password) , bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return nil , err
	}

	user_res.Password = string(hashedPassword)


	user := repository.User{
		Email: user_res.Email,
		Firstname: user_res.Firstname,
		Lastname: user_res.Lastname,
	}

	i_user , err := us.user_repo.CreateUser(user)
	if err != nil {
		logs.Error(err)
	}

	response := UserResponse{
		Email: i_user.Email,
		Firstname: i_user.Firstname,
		Lastname: i_user.Lastname,
	}

	return &response , nil	
}

func (us userService) UpdateUser(id int, user_res UserResponse) (*UserResponse, error)  {
	if id == 0 {
		logs.Error("Invalid ID")
		return nil, errors.New("Invalid ID")
	}

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(user_res.Password) , bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return nil , err
	}

	user_res.Password = string(hashedPassword)

	user := repository.User{
		Email: user_res.Email,
		Firstname: user_res.Firstname,
		Lastname: user_res.Lastname,
	}

	u_user, err := us.user_repo.UpdateUser(id, user)
	if err != nil {
			logs.Error(err)
			return nil, errors.New("Failed to update user")
	}

	response := UserResponse{
		Email: u_user.Email,
		Firstname: u_user.Firstname,
		Lastname: u_user.Lastname,
	}

	return &response, nil
}

func (us userService) DeleteUser(id int) (*UserResponse, error)  {
	if id == 0 {
		logs.Error("Invalid id")
	}

	_ , err := us.user_repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return nil , nil	
}

func (us userService) LoginUser(user_res UserResponse) (*UserResponse, error) {

	user := repository.User{
		Email:    user_res.Email,
		Password: user_res.Password,
	}

	selectedUser, err := us.user_repo.LoginUser(user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user_res.Password))
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Email:     selectedUser.Email,
		Firstname: selectedUser.Firstname,
		Lastname:  selectedUser.Lastname,
	}, nil
}




