package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique; size:50"`
	bio      string `gorm:"size:200; default:null"`
	Password string
}
