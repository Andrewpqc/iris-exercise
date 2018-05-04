package main

import (
	"github.com/kataras/iris"
)

func main(){
	app:=iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello<h2>")
	})
	app.Run(iris.Addr(":8080"),iris.WithConfiguration(iris.Configuration{
		//here is the default configuration
		DisableStartupLog:false,
		DisableInterruptHandler:false,
		DisablePathCorrection:false,
		EnablePathEscape:false,
		FireMethodNotAllowed:false,
		DisableBodyConsumptionOnUnmarshal:false,
		DisableAutoFireStatusCode:false,
		TimeFormat:"Mon,02 Jan 2018 15:04:05 GMT",
		Charset:"UTF_8",
	}))

	//或者在启动之前
	//app.Configure(iris.WithConfiguration(iris.Configuration{
	//	//bala,bala
	//}))
	//app.Run(iris.Addr(":8080"))

}