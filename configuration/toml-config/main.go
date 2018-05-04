//从toml文件中导入配置文件
package main

import (
	"github.com/kataras/iris"
)

func main(){
	app:=iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello,world")
	})

	app.Run(iris.Addr(":8080"),iris.WithConfiguration(iris.TOML("./iris.tml")))

	//or before run
	//app.Configure(iris.WithConfiguration(iris.TOML("./iris.toml")))
	//app.Run(iris.Addr(":8080"))
	}