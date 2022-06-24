package service

import (
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/repository"
	"awesomeProject2/pkg/helper"
	"errors"
)

type LoginService struct {
	userRepository *repository.UserRepository
}

func NewLoginService(userRepository *repository.UserRepository) *LoginService {
	return &LoginService{userRepository: userRepository}
}

func (l *LoginService) Login(loginUser model.User) (*string, error) {
	registeredUser, err := l.userRepository.GetLoginUser(loginUser.Email)
	if err != nil {
		return nil, err
	}
	checked := helper.CheckPasswordMatch(loginUser.Password, registeredUser.Password)
	if checked != nil {
		return nil, errors.New("passwords_do_not_match")
	}

	token, errs := helper.GetToken()
	if errs != nil {
		return nil, errs
	}
	return token, nil
}
