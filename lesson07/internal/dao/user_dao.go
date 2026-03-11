package dao

import (
	"context"
	"lesson07/internal/model"

	"gorm.io/gorm"
)

type UserDAO interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

type userDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{db: db}
}

// CreateUser 创建新用户
func (dao *userDAO) CreateUser(ctx context.Context, user *model.User) error {
	return dao.db.WithContext(ctx).Create(user).Error
}

// GetUserByUsername 查找用户名是否存在(用于登录)
func (dao *userDAO) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}
