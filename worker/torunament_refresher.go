package worker

import (
	"../scraper"
	"fmt"
	"os"
)

func RefreshPremierLeague() {
	data, timeInto, err := scraper.GetFullTeamsInfo(scraper.TournamentB, true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Downloaded data in ", timeInto)
	fmt.Println(data)

	file, err := os.Create("./data/premier_league.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Writing %d bytes", n)
}

