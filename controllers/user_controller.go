package controllers

import (
	"ituring/services"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

func (c *UserController) PostRegister() {
	username := c.Ctx.PostValue("username")
	password := c.Ctx.PostValue("password")
	if username == "" || password == "" {
		c.Ctx.JSON(iris.Map{
			"msg":  "用户名或密码不能为空",
			"code": iris.StatusBadRequest,
		})
		return
	}
	c.Ctx.JSON(iris.Map{
		"msg":  "注册成功！",
		"code": iris.StatusOK,
	})
}

func (c *UserController) PostLogin() {
	username := c.Ctx.PostValue("username")
	password := c.Ctx.PostValue("password")
	if username == "" || password == "" {
		c.Ctx.JSON(iris.Map{
			"msg":  "用户名或密码不能为空",
			"code": iris.StatusBadRequest,
		})
		return
	}
	c.Ctx.JSON(iris.Map{
		"msg":  "登陆成功！",
		"code": iris.StatusOK,
	})
}
