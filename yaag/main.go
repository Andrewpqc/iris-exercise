package main

import (
	"github.com/kataras/iris"
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"fmt"
)

type myXML struct {
	Result string `xml:"result"`
}

func main(){
	app:=iris.New()

	//yaag这个中间件是帮助自动生成api文档的
	//初始化中间件
	yaag.Init(&yaag.Config{
		On: true,
		DocTitle: "Iris",
		DocPath: "apidoc.html",
		BaseUrls:map[string]string{"production":"","staging":""},
	})

	//注册中间件
	app.Use(irisyaag.New())

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"result":"hello,world"})
	})

	app.Get("/plain", func(ctx iris.Context) {
		ctx.Text("hello,world")
	})

	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(myXML{Result:"hello,world"})
	})

	app.Get("/urlParam", func(ctx iris.Context) {
		fmt.Println(ctx.URLParams())
		value:=ctx.URLParam("key")
		ctx.JSON(iris.Map{"value":value})
	})

	app.Run(iris.Addr(":8080"))
}