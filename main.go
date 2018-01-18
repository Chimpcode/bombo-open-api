package main

import (
	"./scraper"
	"fmt"
)

func main() {
	data, err := scraper.GetFullTeamsInfo(false)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}
