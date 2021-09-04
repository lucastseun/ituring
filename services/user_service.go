package services

import (
	"errors"
	"ituring/models"
	"ituring/utils"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (u *UserService) Register(password string, user models.User) error {
	// 检查表是否存在, 没有users表，则创建users表
	if !u.db.Migrator().HasTable(&models.User{}) {
		u.db.Migrator().CreateTable(&models.User{})
	}
	// 检查表里是否已存在该user
	_, hashUser := u.GetUserByNameAndPassword(user.Username, password)
	if hashUser {
		return errors.New("用户已存在，请勿重复注册！")
	}

	hashed, err := models.GeneratePassword(password)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	nanoid, err := utils.NanoId()
	if err != nil {
		return err
	}
	user.AccountId = nanoid

	res := u.db.Create(&user)

	return res.Error
}

func (u *UserService) GetUserByNameAndPassword(username, password string) (models.User, bool) {
	var user models.User

	u.db.Where(&models.User{Username: username}).Find(&user)

	if user.Username == username {
		hashed := user.Password
		if ok, _ := models.ValidatePassword(password, []byte(hashed)); ok {
			return models.User{AccountId: user.AccountId}, true
		}
	}

	return models.User{}, false
}

func (u *UserService) GetUserByID(id string) (bool, models.User) {
	var user models.User

	res := u.db.First(&user, "account_id = ?", id)
	found := false

	if res.RowsAffected > 0 {
		found = true
	}
	return found, user
}

func (u *UserService) DeleteByID(id string) bool {
	var user models.User
	// 永久删除
	res := u.db.Unscoped().Where(&models.User{AccountId: id}).Delete(&user)

	return res.RowsAffected == 1
}
