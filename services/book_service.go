package services

import "ituring/dbsource"

type BookService struct {
}

func NewBookService(db *dbsource.MYSQL) *BookService {
	return &BookService{}
}
