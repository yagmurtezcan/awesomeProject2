package service

import (
	"awesomeProject2/internal/dto"
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/repository"
	"awesomeProject2/pkg/helper"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) GetAllUser() ([]model.User, error) {
	userArray, err := u.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}
	return userArray, nil
}

func (u *UserService) CreateUser(user model.User) (*model.User, error) {
	hashedPassword, errs := helper.HashPassword(user.Password)
	if errs != nil {
		return nil, errs
	}
	user.Password = hashedPassword
	err := u.userRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetUserById(id int) (*model.User, error) {
	user, err := u.userRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) DeleteUser(userId int) error {
	err := u.userRepository.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateUser(userDTO dto.UserDTO) (*model.User, error) {
	user, err := u.GetUserById(userDTO.ID)
	if err != nil {
		return nil, err
	}

	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Email = userDTO.Email

	updatedUser, updateErr := u.userRepository.UpdateUser(user)
	if updateErr != nil {
		return nil, updateErr
	}
	return updatedUser, nil
}
