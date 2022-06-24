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

func (u *UserRepository) GetAllUser() ([]model.User, error) {
	var user []model.User
	if err := u.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) CreateUser(user *model.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) DeleteUser(userId int) error {
	var user model.User
	if err := u.db.Where("id=?", userId).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	if err := u.db.Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetLoginUser(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
