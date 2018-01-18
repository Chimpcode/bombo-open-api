package main

import (
	"./scraper"
	"fmt"
)

func main() {
	data, err := scraper.GetFullTeamsInfo()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}
