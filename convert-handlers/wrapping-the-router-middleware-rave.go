package main

/**
以中间件的形式将raven服务集成到应用中
*/
//怎样写一个中间件

import (
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"runtime/debug"
	"fmt"
	"errors"
)

//自定义一个中间件
func irisRavenMiddleware(ctx iris.Context) {
	w, r := ctx.ResponseWriter(), ctx.Request()

	defer func() {
		if rval := recover(); rval != nil {
			debug.PrintStack()
			rvalStr := fmt.Sprint(rval)
			packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
			raven.Capture(packet, nil)
			w.WriteHeader(iris.StatusInternalServerError)
		}
	}()

	ctx.Next()
}

func init(){
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}
func main(){
	app:=iris.New()
	app.Use(irisRavenMiddleware)

	app.Get("/hello", func(ctx iris.Context) {
		ctx.Text("hello,world")
	})

	app.Run(iris.Addr(":8080"),iris.WithoutVersionChecker)
}