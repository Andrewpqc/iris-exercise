//自定义Unmarshal函数，来将请求体序列化到go结构

package main

import (
	"github.com/kataras/iris"
	"gopkg.in/yaml.v2"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Post("/", handler)
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8000"), iris.WithoutVersionChecker, iris.WithOptimizations)
}

type config struct {
	Addr       string `yaml:"addr"`
	ServerName string `yaml:"serverName"`
}

func (c *config) Decode(body []byte) error {
	return yaml.Unmarshal(body, c)
}

func handler(ctx iris.Context) {
	var c config

	//读取请求体的内容
	if err := ctx.UnmarshalBody(&c, nil); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v", c)
}
