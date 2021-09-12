package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountId string `gorm:"account_id" json:"accountId"`
	Username  string `gorm:"username" json:"username"`
	Password  string `gorm:"password" json:"password"`
}

type UserClaims struct {
	Username  string
	AccountId string
}
