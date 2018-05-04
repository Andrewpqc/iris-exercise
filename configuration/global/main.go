//共享家目录中的配置文件，可以达到多个iris实例共享同一份配置文件
//配置文件必须存在与家目录的iris.yaml
package main

import (
	"github.com/kataras/iris"
)

func main(){
	app:=iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("hello,world")
	})


	// Good when you share configuration between multiple iris instances.
	// This configuration file lives in your $HOME/iris.yml for unix hosts
	// or %HOMEDRIVE%+%HOMEPATH%/iris.yml for windows hosts, and you can modify it.
	app.Run(iris.Addr(":8080"),iris.WithGlobalConfiguration)
}