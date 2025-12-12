package api

import (
	"net/http"
	"time"

	"lesson07/dao"
	"lesson07/model"
	"lesson07/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 如果用户存在，这里这种是用户名可以一致的，即只要密码不一致就视为不同用户
	if dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 检查用户是否存在且密码是否正确
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	// 生成jwt token
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	// 返回token
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
		"token":   token,
	})
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ChangePassword(c *gin.Context) {
	// 注意：这个函数被 user/password 调用，而该路由在 router.go 中已经过了 JWTAuth 中间件
	// 所以，我们可以直接从上下文中拿到 **已经验证过的** 用户名
	currentUsername, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "无法获取用户信息"})
		return
	}

	// 1. 定义请求体结构
	type ChangePasswordReq struct {
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}
	var req ChangePasswordReq

	// 2. 绑定JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
		return
	}

	// 3. 调用DAO层进行密码修改
	success := dao.ChangeUserPassword(currentUsername.(string), req.OldPassword, req.NewPassword)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "旧密码错误"})
		return
	}

	// 4. 返回成功
	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}
