//还是不太明白这个中间件是干什么的

package main

import (
	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/newrelic"
)

func main() {
	app := iris.New()
	//licenseKey需要40个字符长度
	config := newrelic.Config("APP_SERVER_NAME", "0123456789012345678901234567890123456789")
	config.Enabled = true
	m, err := newrelic.New(config)
	if err != nil {
		app.Logger().Fatal(err)
	}
	app.Use(m.ServeHTTP)

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("success!\n")
	})

	app.Run(iris.Addr(":8080"))
}