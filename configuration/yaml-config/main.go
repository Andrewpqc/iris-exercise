package main

import (
	"github.com/kataras/iris"
)

func main(){
	app:=iris.New()
	app.Get("/", func(ctx iris.Context) {
			ctx.WriteString("hello,world")
	})

	app.Run(iris.Addr(":8080"),iris.WithConfiguration(iris.YAML("./iris.yml")))

	// or before run:
	// app.Configure(iris.WithConfiguration(iris.YAML("./configs/iris.yml")))
	// app.Run(iris.Addr(":8080"))
}