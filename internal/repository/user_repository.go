package repository

import (
	"awesomeProject2/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetAllUser() []model.User {
	var user []model.User
	u.db.Find(&user)
	return user
}

func (u *UserRepository) CreateUser(user model.User) (*model.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
