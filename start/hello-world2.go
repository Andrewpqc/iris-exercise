package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main(){
	app:=iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET","/", func(ctx iris.Context) {
		ctx.HTML("<h2>hello,world 2 by iris<h2>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"hello":"world"})
	})

	app.Run(iris.Addr("127.0.0.1:8000"),iris.WithoutVersionChecker,iris.WithoutServerError(iris.ErrServerClosed))
}

