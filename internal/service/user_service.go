package service

import (
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) GetAllUser() []model.User {
	return u.userRepository.GetAllUser()
}

func (u *UserService) CreateUser(user model.User) (*model.User, error) {
	createUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}
