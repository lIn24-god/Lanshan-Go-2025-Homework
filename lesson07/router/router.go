package router

import (
	"lesson07/internal/handler"
	"lesson07/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userHandler *handler.UserHandler
}

func NewRouter(userHandler *handler.UserHandler) *Router {
	return &Router{userHandler: userHandler}
}

func (r *Router) SetUp(engine *gin.Engine) {

	authMiddleware := middleware.AuthMiddleware()

	//公共路由
	public := engine.Group("/api")
	{
		public.POST("/user/login", r.userHandler.Login)
		public.POST("/user/register", r.userHandler.Register)
	}

	//需要认证的路由
	protected := engine.Group("/api")
	protected.Use(authMiddleware)
	{

	}
}
