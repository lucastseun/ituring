package main

import (
	"ituring/routes"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	routes.InitRouter(app)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
