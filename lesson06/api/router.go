package api

import (
	"github.com/gin-gonic/gin"
	"lesson06/middleware"
)

func InitRouter_gin() {
	r := gin.Default()
	r.GET("/ping", middleware.Example1(), middleware.Example2(), Ping1)
	r.POST("login", Login)
	r.Run(":8080")
}
