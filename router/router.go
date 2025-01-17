package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-learn/controller"
	"iris-learn/database"
	"iris-learn/environment"
	"iris-learn/service"
)

// NewRouter creates a new iris application with the routes and controllers
func NewRouter() *iris.Application {
	app := iris.New()
	app.Get("/ping", pong).Describe("ping pong")

	mvc.Configure(app.Party("/greet"), setup)

	return app
}

func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}

func setup(app *mvc.Application) {
	app.Register(environment.DEV,
		database.NewDB,
		service.NewGreetService,
	)

	app.Handle(new(controller.GreetController))
}
