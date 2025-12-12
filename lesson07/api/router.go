package api

import (
	"lesson07/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterGin() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS()) //全局中间件
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.GET("/ping", pong)
	r.PUT("/user/password", middleware.JWTAuth(), ChangePassword) //在修改密码前先验证身份

	return r
}
