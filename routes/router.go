package routes

import (
	"ituring/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	user := app.Party("/user")
	mvc.New(user).Handle(new(controllers.UserController))
}
