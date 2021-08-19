package services

import "ituring/dbsource"

type UserService struct {
}

func NewUserService(db *dbsource.MYSQL) *UserService {
	return &UserService{}
}
