package main

import (
	"os"
	"encoding/json"
	"strings"
	"log"
	"./scraper"
)

func init() {
	if os.IsExist(os.Mkdir("./data", 0755)) {
		log.Println("Data directory already exist")
	}
}


func normalizeName(name string) string {
	return strings.Replace(strings.ToLower(name), " ", "_", -1)
}

func SaveDataForMatchWork(w *Work) error {
	matchEvents, err :=scraper.GetEventsFromMatch(w.URL)
	if err != nil {
		return err
	}

	file, err := os.Create("./data/match_" + normalizeName(w.Name) + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(matchEvents, "", "	")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func SaveDataForLeagueWork(w *Work) error {
	teams, err := scraper.GetFullTeamsInfo(w.URL, true)
	if err != nil {
		return err
	}


	file, err := os.Create("./data/league_" + normalizeName(w.Name) + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(teams, "", "	")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil

}
