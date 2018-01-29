package main

import (
	"fmt"

	"./scraper"
)

func main() {
	//app := iris.Default()
	//worker.RefreshPremierLeague()
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowCredentials: true,
	//})
	//
	//api := app.Party("/api/v1")
	//api.Use(crs)
	//server.LinkApi(api)
	//
	//app.Run(iris.Addr(":9800"))

	data, err := scraper.GetEventsFromMatch(scraper.ExampleMatch)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)

}
