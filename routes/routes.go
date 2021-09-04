package routes

import (
	"ituring/handlers"
	"ituring/services"
	"ituring/utils"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func InitRouter(app *iris.Application, db *gorm.DB) {
	var (
		userService = services.NewUserService(db)
		bookService = services.NewBookService(db)
	)
	// http://localhost/user
	user := app.Party("/user")
	user.Use(utils.VerifyMiddleware())
	{
		handler := handlers.NewUserHandler(userService)
		user.Post("/register", handler.Register)
		user.Post("/login", handler.Login)
		user.Post("/delete", handler.Delete)
	}
	// http://localhost/book
	book := app.Party("book")
	{
		handler := handlers.NewBookHandler(bookService)
		book.Post("/hot", handler.Hot)
	}
}
