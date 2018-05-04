//sentry服务在iris中的集成
//貌似这个sentry服务没有什么卵用,手机错误信息的
//文档参考:https://blog.windrunner.me/tool/sentry.html
package main

import (
	"github.com/kataras/iris"
	"github.com/getsentry/raven-go"

	"runtime/debug"
	"net/http"
	"fmt"
	"errors"
)

func init(){
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")

}

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hi")
	})

	// Example for WrapRouter is already here:
	// https://github.com/kataras/iris/blob/master/_examples/routing/custom-wrapper/main.go#L53
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, irisRouter http.HandlerFunc) {
		// Exactly the same source code:
		// https://github.com/getsentry/raven-go/blob/379f8d0a68ca237cf8893a1cdfd4f574125e2c51/http.go#L70

		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		irisRouter(w, r)
	})

	app.Run(iris.Addr(":8080"))
}