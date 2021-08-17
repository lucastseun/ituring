package main

import (
	"fmt"
	"ituring/models"
	"ituring/routes"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "root:123456@tcp(127.0.0.1:3306)/ituring?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	if !db.Migrator().HasTable(&models.User{}) {
		db.Migrator().CreateTable(&models.User{})
	}
	app.Listen(":3000")
}
