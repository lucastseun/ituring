package handlers

import (
	"fmt"
	"ituring/models"
	"ituring/services"

	"github.com/kataras/iris/v12"
)

type UserHandler struct {
	service *services.UserService
}

type Response struct {
	AccountId string `json:"accountId"`
	Token     string `json:"token"`
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
	u := h.service.Register(password, models.User{
		Username: username,
	})

	if u != nil {
		ctx.JSON(iris.Map{
			"errMsg": fmt.Sprintf("%s", u),
			"code":   iris.StatusOK,
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
			"errMsg": "用户名或密码不能为空!",
			"code":   iris.StatusOK,
		})
		return
	}
	user, found := h.service.GetUserByNameAndPassword(username, password)
	if !found {
		ctx.JSON(iris.Map{
			"errMsg": "用户未注册！",
			"code":   iris.StatusOK,
		})
		return
	}

	ctx.JSON(iris.Map{
		"msg": "登陆成功！",
		"data": Response{
			AccountId: user.AccountId,
			Token:     "",
		},
		"code": iris.StatusOK,
	})
}

func (h *UserHandler) Delete(ctx iris.Context) {
	accountId := ctx.PostValue("accountId")

	if accountId == "" {
		ctx.JSON(iris.Map{
			"errMsg": "必填参数为空!",
			"code":   iris.StatusBadRequest,
		})
		return
	}

	logoff := h.service.DeleteByID(accountId)
	if logoff {
		ctx.JSON(iris.Map{
			"msg":  "注销成功！",
			"code": iris.StatusOK,
		})
		return
	}
	ctx.JSON(iris.Map{
		"errMsg": "注销失败！",
		"code":   iris.StatusOK,
	})
}
