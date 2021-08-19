package main

import (
	"ituring/dbsource"
	"ituring/routes"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	db, err := dbsource.ConnectMYSQL()
	if err != nil {
		log.Fatalf("error connecting to the MySQL database: %v", err)
	}
	routes.InitRouter(app, db)
	return app
}

func main() {
	app := newApp()
	app.Listen(":3000")
}
