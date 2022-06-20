package repository

import (
	"awesomeProject2/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetAllUser() []entity.User {
	var user []entity.User
	u.db.Find(&user)
	return user
}
