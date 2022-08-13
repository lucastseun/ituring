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
		userService      = services.NewUserService(db)
		bookService      = services.NewBookService(db)
		verifyMiddleware = utils.VerifyMiddleware()
	)
	// http://localhost/user
	user := app.Party("/user")
	{
		handler := handlers.NewUserHandler(userService)
		user.Post("/register", handler.Register)
		user.Post("/signin", handler.Login)
		user.Post("/delete", verifyMiddleware, handler.Delete)
	}
	// http://localhost/book
	book := app.Party("book")
	{
		book.Use(verifyMiddleware)
		handler := handlers.NewBookHandler(bookService)
		book.Post("/hot", handler.Hot)
	}
}
