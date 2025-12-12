package dao

import (
	"gorm.io/gorm"
	"lesson07/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Create(user *model.User) error {
	return dao.db.Create(user).Error
}
