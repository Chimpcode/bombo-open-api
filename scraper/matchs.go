package scraper

import (
	"github.com/anaskhan96/soup"
	"encoding/json"
)

func getScoreFromMatch(urlMatch string) ([]byte, error) {
	resp, err := soup.Get(urlMatch)

	if err != nil {
		return []byte{}, err
	}

	doc := soup.HTMLParse(resp)

	pResult := doc.Find("td", "class", "current-result")

	scoreHome := ""
	scoreAway := ""
	teamHome := ""
	teamAway := ""
	for i, c := range pResult.FindAll("span") {
		if i == 0 {
			scoreHome = c.Text()
		}else if i == 2 {
			scoreAway = c.Text()
		}
	}

	tHome := doc.Find("td", "class", "tname-home")
	teamHome = tHome.Find("span").Find("a").Text()

	tAway := doc.Find("td", "class", "tname-away")
	teamAway = tAway.Find("span").Find("a").Text()

	finalData := map[string]map[string]string{}
	finalData["home"] = map[string]string{teamHome:scoreHome}
	finalData["away"] = map[string]string{teamAway:scoreAway}

	return json.Marshal(finalData)
}
