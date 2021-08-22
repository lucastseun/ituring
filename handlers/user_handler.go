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
	if h.service.RegisterByNameAndPassword(username, password) {
		ctx.JSON(iris.Map{
			"msg":  "注册成功！",
			"code": iris.StatusOK,
		})
		return
	}
	ctx.JSON(iris.Map{
		"msg":  "注册失败！",
		"code": iris.StatusOK,
	})
}

func (h *UserHandler) Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username == "" || password == "" {
		ctx.JSON(iris.Map{
			"msg":  "用户名或密码不能为空!",
			"code": iris.StatusBadRequest,
		})
		return
	}
	u, found := h.service.GetUserByNameAndPassword(username, password)
	if !found {
		ctx.JSON(iris.Map{
			"msg":  "用户未注册！",
			"code": iris.StatusOK,
		})
		return
	}
	ctx.JSON(iris.Map{
		"msg":  "登陆成功！",
		"data": u,
		"code": iris.StatusOK,
	})
}

func (h *UserHandler) Delete(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username == "" || password == "" {
		ctx.JSON(iris.Map{
			"msg":  "用户名或密码不能为空!",
			"code": iris.StatusBadRequest,
		})
		return
	}
	ctx.JSON(iris.Map{
		"msg":  "删除成功！",
		"code": iris.StatusOK,
	})
}
