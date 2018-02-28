package main

import (
	"time"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
)

func main() {

	manager, err:= LoadManagerFromFile("./state/manager_state.json")
	if err != nil {
		panic(err)
	}

	manager.Run(1*time.Second)

	app := iris.Default()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	app.Use(crs)

	api := app.Party("/api/v1.0/")

	LinkApi(api, manager)

	app.Run(iris.Addr(":8500"))
}

