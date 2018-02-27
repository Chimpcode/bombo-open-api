package main

import (
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

	//data, err := scraper.GetEventsFromMatch(scraper.ExampleMatch)
	//if err != nil {
	//	panic(err.Error())
	//}
	//pp.Println(data)

	teams, _, _:= scraper.GetFullTeamsInfo("https://www.scoreboard.com/es/futbol/inglaterra/premier-league/equipos", true)
	CreateCSVFromStruct(teams)


	//manager := NewWorkManager()
	//work := NewWork(
	//	"https://www.scoreboard.com/es/futbol/inglaterra/premier-league/equipos",
	//	"premier_league",
	//	WorkLeagueType,
	//	30*time.Second,
	//)
	//
	//workMatch := NewWork(
	//	"https://www.scoreboard.com/es/partido/eibar-malaga-cf-2017-2018/Amqd6uGr",
	//	"premier_league",
	//	WorkMatchType,
	//	20*time.Second,
	//)
	//
	//manager.AddWork(work)
	//manager.AddWork(workMatch)
	//
	//
	//manager.Run(1*time.Second)
	//
	//for {
	//	h, m, s := time.Now().Clock()
	//	fmt.Printf("%d:%d:%d\n", h, m, s)
	//	time.Sleep(1 * time.Second)
	//}


}
