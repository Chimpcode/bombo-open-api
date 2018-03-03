package main

import (
	"fmt"
	"bytes"
	"os"
	"./scraper"
)

func CreateCSVFromStruct(teams map[string]*scraper.Team) {
	buffer := bytes.NewBufferString("")
	for _, team := range teams {

		for _, player := range team.GoalKeeper {
			s := fmt.Sprintf("%s,%s,%d\n", player.Name, player.Team, player.Cost)
			buffer.Write([]byte(s))
		}
		for _, player := range team.Defender {
			s := fmt.Sprintf("%s,%s,%d\n", player.Name, player.Team, player.Cost)
			buffer.Write([]byte(s))
		}
		for _, player := range team.MidFielder {
			s := fmt.Sprintf("%s,%s,%d\n", player.Name, player.Team, player.Cost)
			buffer.Write([]byte(s))
		}
		for _, player := range team.Forwarder {
			s := fmt.Sprintf("%s,%s,%d\n", player.Name, player.Team, player.Cost)
			buffer.Write([]byte(s))
		}
	}

	file, err := os.Create("./data/league_csv.csv")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
}

