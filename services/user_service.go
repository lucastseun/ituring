package services

import (
	"ituring/models"

	"gorm.io/gorm"
)

type UserService struct {
	Service *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{Service: db}
}

func (u *UserService) RegisterByNameAndPassword(username, password string) bool {
	user := models.User{
		Username: username,
		Password: password,
	}
	_, found := u.GetUserByNameAndPassword(username, password)
	if found {
		return false
	}
	res := u.Service.Create(&user)
	return res.Error == nil
}

func (u *UserService) GetUserByNameAndPassword(username, password string) (models.User, bool) {
	var user models.User

	res := u.Service.Where(&models.User{Username: username, Password: password}).Find(&user)
	found := false

	if res.RowsAffected > 0 {
		found = true
	}
	return user, found
}

func (u *UserService) GetUserByID(id uint64) (models.User, bool) {
	var user models.User

	res := u.Service.First(&user, "id = ?", id)
	found := false

	if res.RowsAffected > 0 {
		found = true
	}
	return user, found
}

func (u *UserService) DeleteUserByNameAndPassword(username, password string) (models.User, bool) {
	var user models.User

	res := u.Service.Unscoped().Where(&models.User{Username: username, Password: password}).Delete(&user)
	if res.Error != nil {
		return user, false
	}
	return user, true
}
