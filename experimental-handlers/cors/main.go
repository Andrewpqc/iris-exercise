//cors设定允许的origins字段
//origins字段指定请求来自于哪个站点，指示请求的服务器名称
//Access-Control-Allow-Origin: foo.com用于实现服务器端的访问控制

package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
)

func main(){
	app:=iris.New()

	//这里如果客户端不带Origin请求头的话好像也可以直接请求导资源
	crs:=cors.New(cors.Options{
		AllowedOrigins:[]string{"aaa.com"},//// allows everything, use that to change the hosts.
		AllowCredentials:true,

	})

	v1:=app.Party("/api/v1",crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/home", func(ctx iris.Context) {
			ctx.WriteString("Hello from /home")
		})
		v1.Get("/about", func(ctx iris.Context) {
			ctx.WriteString("Hello from /about")
		})
		v1.Post("/send", func(ctx iris.Context) {
			ctx.WriteString("sent")
		})
		v1.Put("/send", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/send", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
	}

	app.Run(iris.Addr(":8080"),iris.WithoutVersionChecker)
}