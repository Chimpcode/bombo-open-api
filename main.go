package main

import (
	"time"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
)

func main() {

	//events, err := scraper.GetEventsFromMatch("https://www.scoreboard.com/es/partido/brighton-swansea-2017-2018/rTJMoCnn/")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//pp.Println(events)
	manager, err:= LoadManagerFromFile("./state/manager_state.json")
	if err != nil {
		panic(err)
	}

	manager.Run(1*time.Second)

	app := iris.Default()

	crs := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	crsConf := cors.NewAllowAllAppMiddleware()

	app.Configure(crsConf)
	app.Use(crs)

	api := app.Party("/api/v1.0/")
	managerApi := app.Party("/manager/")


	//bauth := basicauth.New(basicauth.Config{
	//	Users: map[string]string{
	//		"bregymr": "malpartida1",
	//	},
	//	Expires: 24*time.Hour,
	//	Realm: "bombo",
	//})

	//managerApi.Use(bauth)

	LinkApi(api, manager)

	LinkManagerApi(managerApi, manager)

	app.Logger().SetLevel("debug")



	app.Run(iris.Addr(":8500"))
}

