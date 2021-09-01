package services

import (
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

func (u *UserService) RegisterByNameAndPassword(username, password string) bool {
	user := models.User{
		Username: username,
		Password: password,
	}
	_, found := u.GetUserByNameAndPassword(username, password)
	if found {
		return false
	}
	id, err := utils.NanoId()
	if err != nil {
		panic(err)
	}
	user.AccountId = id
	res := u.db.Create(&user)
	return res.Error == nil
}

func (u *UserService) GetUserByNameAndPassword(username, password string) (models.User, bool) {
	var user models.User

	res := u.db.Where(&models.User{Username: username, Password: password}).Find(&user)
	found := false

	if res.RowsAffected > 0 {
		found = true
	}
	return user, found
}

func (u *UserService) GetUserByID(id uint64) (models.User, bool) {
	var user models.User

	res := u.db.First(&user, "id = ?", id)
	found := false

	if res.RowsAffected > 0 {
		found = true
	}
	return user, found
}

func (u *UserService) DeleteUserByNameAndPassword(username, password string) (models.User, bool) {
	var user models.User

	res := u.db.Unscoped().Where(&models.User{Username: username, Password: password}).Delete(&user)
	if res.Error != nil {
		return user, false
	}
	return user, true
}
