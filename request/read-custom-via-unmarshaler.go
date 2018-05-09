//使用内置的Unmarshal函数，来将请求体序列化到go结构

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gopkg.in/yaml.v2"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Post("/", handler)

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}


type config struct {
	Addr       string `yaml:"addr"`
	ServerName string `yaml:"serverName"`
}


func handler(ctx iris.Context) {
	var c config

	//使用yaml内部自定义的Unmarshal函数将请求体反序列化到结构
	if err := ctx.UnmarshalBody(&c, context.UnmarshalerFunc(yaml.Unmarshal)); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v", c)
}