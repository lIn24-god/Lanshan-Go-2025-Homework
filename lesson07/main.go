package main

import (
	"lesson07/api"
	"lesson07/dao"
	"lesson07/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/lesson07?charset=utf8mb4&parseTime=True&loc=local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})

	userDAO := dao.NewUserDao(db)

	r := api.InitRouterGin()

	r.Run(":8080")
}
