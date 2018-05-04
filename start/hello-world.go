package main

import (
	"github.com/kataras/iris"
)

func main(){
	//without log
	//app:=iris.New()

	//with log
	app:=iris.Default()
	app.Handle("GET","/hello", func(ctx iris.Context) {
		ctx.HTML("<h1>hello,world!</h1>")
	})
	app.Handle("GET","/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"hello":"world","arg2":1})
	})

	app.Handle("GET","/str", func(ctx iris.Context) {
		ctx.Write([]byte("here is a string send you by iris"))
	})

	app.Run(iris.Addr("127.0.0.1:8080"),iris.WithoutVersionChecker)
}
