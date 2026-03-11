package main

import (
	"fmt"
	"lesson07/internal/dao"
	"lesson07/internal/handler"
	"lesson07/internal/model"
	"lesson07/internal/service"
	"lesson07/router"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(localhost:3306)/lesson07?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		fmt.Println("Failed to run migrate:", err)
		return
	}

	userDAO := dao.NewUserDAO(db)
	userService := service.NewUserService(userDAO)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	routerInstance := router.NewRouter(userHandler)
	routerInstance.SetUp(r)

	//启动gin服务
	err1 := r.Run(":8080")
	if err1 != nil {
		return
	}
}
