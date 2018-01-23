package main

import (
	"github.com/kataras/iris"
	"./server"
)

func main() {
	app := iris.Default()
	api := app.Party("/api/v1")
	server.LinkApi(api)

	app.Run(iris.Addr(":9800"))
}
