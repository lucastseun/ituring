package handlers

import (
	"ituring/services"

	"github.com/kataras/iris/v12"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service}
}

func (h *BookHandler) Hot(ctx iris.Context) {

}
