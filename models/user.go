package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"name" json:"username"`
	Password string `gorm:"password" json:"password"`
}
