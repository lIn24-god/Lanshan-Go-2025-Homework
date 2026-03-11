package service

import (
	"context"
	"lesson07/internal/dao"
	"lesson07/internal/model"
	"lesson07/pkg/encrypt"
	"lesson07/pkg/jwt"
)

type UserService interface {
	RegisterUser(ctx context.Context, username, password string) (*model.User, error)
	LoginUser(ctx context.Context, username, password string) (string, *model.User, error)
}

type userService struct {
	userDAO dao.UserDAO
}

func NewUserService(userDAO dao.UserDAO) UserService {
	return &userService{userDAO: userDAO}
}

func (s *userService) RegisterUser(ctx context.Context, username, password string) (*model.User, error) {

	hashPassword, _ := encrypt.HashPassword(password)

	user := &model.User{
		Username: username,
		Password: hashPassword,
	}

	if err := s.userDAO.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) LoginUser(ctx context.Context, username, password string) (string, *model.User, error) {
	user, err := s.userDAO.GetUserByUsername(ctx, username)
	if err != nil {
		return "", nil, err
	}

	if !encrypt.CheckHashPassword(password, user.Password) {
		return "", nil, err
	}

	// 3. 生成 JWT token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err // token 生成失败，系统错误
	}

	return token, user, nil
}
