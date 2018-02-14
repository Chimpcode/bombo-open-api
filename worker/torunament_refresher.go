package worker

import (
	"fmt"
	"os"
	"github.com/Chimpcode/bombo-open-api/scraper"
)

func RefreshPremierLeague() {
	data, timeInto, err := scraper.GetFullTeamsInfo(scraper.TournamentB, true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Downloaded data in ", timeInto)
	fmt.Println(data)

	err = os.Mkdir("./data", os.ModeDir)
	if err != nil {
		panic(err.Error())
	}
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

