package routes

import (
	"ituring/dbsource"
	"ituring/handlers"
	"ituring/services"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application, db *dbsource.MYSQL) {
	var (
		userService = services.NewUserService(db)
		bookService = services.NewBookService(db)
	)
	// http://localhost/user
	user := app.Party("/user")
	{
		handler := handlers.NewUserHandler(userService)
		user.Post("/register", handler.Register)
		user.Post("/login", handler.Login)
	}
	// http://localhost/book
	book := app.Party("book")
	{
		handler := handlers.NewBookHandler(bookService)
		book.Post("/hot", handler.Hot)
	}
}
