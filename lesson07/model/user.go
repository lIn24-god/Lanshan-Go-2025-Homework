package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}
