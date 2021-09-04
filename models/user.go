package models

import (
	"golang.org/x/crypto/bcrypt"
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

// 生成哈希值
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// 校验密码是否匹配
func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
