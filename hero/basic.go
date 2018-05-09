//hero这个包要好好学一学，感觉对于项目的组织会有很大的帮助


package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)
func hello(to string) string {
	return "Hello " + to
}

func main(){
	app:=iris.New()
	helloHandler:=hero.Handler(hello)
	app.Get("/{to:string}/hello",helloHandler)
	app.Run(iris.Addr(":8000"),iris.WithoutVersionChecker)
}