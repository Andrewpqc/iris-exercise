package main

import (
	"github.com/kataras/iris"
	"net/http"
	"fmt"
)

//类似于python中的装饰器,在所有的handler之前进行了拦截
func negronilikeTestMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == "/ok" && r.Method == "GET" {
		w.Write([]byte("OK. "))
		next(w, r) // go to the next route's handler
		return
	}
	// else print an error and do not forward to the route's handler.
	w.WriteHeader(iris.StatusBadRequest)
	w.Write([]byte("Bad request"))
}
//没有next,组织了处理器函数往下执行
func nativeTestMiddleware(w http.ResponseWriter,r *http.Request){
	fmt.Println("request url:",r.URL)
}
func main (){
	app:=iris.New()

	//以中间件的形式，会对所有的handler都使用上
	irisMiddle:=iris.FromStd(negronilikeTestMiddleware)
	nativeirisMiddleware:=iris.FromStd(nativeTestMiddleware)
	app.Use(irisMiddle)
	app.Use(nativeirisMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Home </h1>")
		// this will print an error,
		// this route's handler will never be executed because the middleware's criteria not passed.
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx iris.Context) {
		ctx.Writef("Hello world!")
		// this will print "OK. Hello world!".
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(iris.Addr(":8080"))

}