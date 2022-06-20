package service

import (
	"awesomeProject2/internal/entity"
	"awesomeProject2/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) GetAllUser() []entity.User {
	return u.userRepository.GetAllUser()
}
