package main

import (
	"time"
)

func main() {


	//
	//app := iris.Default()
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowCredentials: true,
	//})

	manager := NewWorkManager()
	work := NewWork(
		"https://www.scoreboard.com/es/futbol/inglaterra/premier-league/equipos",
		"premier_league",
		WorkLeagueType,
		60*time.Minute,
	)

	workMatch := NewWork(
		"https://www.scoreboard.com/es/partido/eibar-malaga-cf-2017-2018/Amqd6uGr",
		"premier_league",
		WorkMatchType,
		5*time.Minute,
	)

	manager.AddWork(work)
	manager.AddWork(workMatch)

	err := SaveWorkManagerState(manager)
	if err != nil {
		panic(err)
	}
	//manager.Run(1*time.Second)
	//
	//for {
	//	h, m, s := time.Now().Clock()
	//	fmt.Printf("%d:%d:%d\n", h, m, s)
	//	time.Sleep(1 * time.Second)
	//}


}
