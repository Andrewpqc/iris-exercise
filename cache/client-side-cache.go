package main

import (
	"time"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris"
)

const refreshEvery = 10 * time.Second

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(iris.Cache304(refreshEvery))
	// same as:
	// app.Use(func(ctx iris.Context) {
	// 	now := time.Now()
	// 	if modified, err := ctx.CheckIfModifiedSince(now.Add(-refresh)); !modified && err == nil {
	// 		ctx.WriteNotModified()
	// 		return
	// 	}

	// 	ctx.SetLastModified(now)

	// 	ctx.Next()
	// })

	app.Get("/", greet)
	app.Run(iris.Addr(":8080"))
}

func greet(ctx iris.Context) {
	ctx.Header("X-Custom", "my  custom header")
	ctx.Writef("Hello World! %s", time.Now())
}