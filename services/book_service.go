package services

import "gorm.io/gorm"

type BookService struct {
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{}
}
