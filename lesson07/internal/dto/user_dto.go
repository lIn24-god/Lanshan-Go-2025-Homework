package dto

// UserBrief 用户简要信息（可复用）
type UserBrief struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录相应
type LoginResponse struct {
	Token string    `json:"token"`
	User  UserBrief `json:"user"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterResponse 注册响应
type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
