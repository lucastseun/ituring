package services

import "ituring/models"

type UserService interface {
	GetAll() []models.User
	GetByID(id int) (models.User, bool)
	GetByUsernameAndPasswrod(username, password string) (models.User, bool)
	DeleteByID(id int) bool

	Update(id int, user models.User) (models.User, error)
	UpdatePassword(id int64, newPassword string) (models.User, error)
	UpdateUsername(id int64, newUsername string) (models.User, error)

	Create(userPassword string, user models.User) (models.User, error)
}

type userService struct {
}
