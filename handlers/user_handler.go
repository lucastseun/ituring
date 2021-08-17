package handlers

import (
	"ituring/services"

	"github.com/kataras/iris/v12"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Register(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username == "" || password == "" {
		ctx.JSON(iris.Map{
			"msg":  "用户名或密码不能为空",
			"code": iris.StatusBadRequest,
		})
		return
	}
	ctx.JSON(iris.Map{
		"msg":  "注册成功！",
		"code": iris.StatusOK,
	})
}

func (h *UserHandler) Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username == "" || password == "" {
		ctx.JSON(iris.Map{
			"msg":  "用户名或密码不能为空",
			"code": iris.StatusBadRequest,
		})
		return
	}
	ctx.JSON(iris.Map{
		"msg":  "登陆成功！",
		"code": iris.StatusOK,
	})
}
