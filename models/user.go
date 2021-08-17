package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"name" json:"name"`
	Age      int    `gorm:"age" json:"age"`
	Password string `gorm:"password" json:"password"`
}
