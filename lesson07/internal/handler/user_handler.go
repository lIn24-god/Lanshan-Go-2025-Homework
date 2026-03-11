package handler

import (
	"lesson07/internal/dto"
	"lesson07/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService // 依赖 Service 接口
}

// NewUserHandler 构造函数
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register 用户注册
func (h *UserHandler) Register(c *gin.Context) {
	var request = dto.LoginRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx := c.Request.Context()

	user, err := h.userService.RegisterUser(ctx, request.Username, request.Password)
	if err != nil {
		// 使用全局错误处理函数
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register"})
		return
	}

	resp := dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, resp)
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var request = dto.LoginRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx := c.Request.Context()
	token, user, err := h.userService.LoginUser(ctx, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		return
	}

	resp := dto.LoginResponse{
		Token: token,
		User: dto.UserBrief{
			ID:       user.ID,
			Username: user.Username,
		},
	}

	c.JSON(http.StatusOK, resp)
}
