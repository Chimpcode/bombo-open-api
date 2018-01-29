package scraper

import (
	"github.com/anaskhan96/soup"
	"encoding/json"
	"github.com/gocolly/colly"
	"fmt"
)

func GetScoreFromMatch(urlMatch string) ([]byte, error) {
	//
	//scoreHome := ""
	//scoreAway := ""
	teamHome := ""
	teamAway := ""

	c := colly.NewCollector()

	c.OnHTML("td#flashscore_column", func(e *colly.HTMLElement) {
		teamHome = e.DOM.Find("td.tname-home").Find("a").Text()
		teamAway = e.DOM.Find("td.tname-away").Find("a").Text()
		fmt.Println("On")
		fmt.Println(teamHome, teamAway)
	})

	c.OnHTML("td.current-result", func(e *colly.HTMLElement) {
		spans := e.DOM.Find("span")
		if spans.Length() != 3 {return}
		score := spans.Text()
		fmt.Println(score)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	c.Visit(urlMatch)

	return []byte{}, nil

}


func GetEventsFromMatch(urlMatch string) ([]byte, error) {
	resp, err := soup.Get(urlMatch)
	if err != nil {
		return []byte{}, err
	}

	doc := soup.HTMLParse(resp)

	parts := doc.Find("table", "class", "parts")
	bodyParts := parts.Find("tbody") // REPAIR!!

	currentTitleEvents := ""

	events := map[string][]interface{}{}

	for _, col := range bodyParts.FindAll("tr") {
		// col.Attrs()["class"] == "odd" || col.Attrs()["class"] == "even"

		if len(col.Attrs()) > 0 {
			for _, player := range col.FindAll("td") {
				chunkName := player.Find("div", "class", "name").Find("a")
				playerName := chunkName.Text()
				playerHref := chunkName.Attrs()["href"]

				eventType := "NaN"
				eventDescription := ""
				lineup := player.Find("div", "class", "icon-lineup")
				if lineup.Pointer !=  nil {
					if lineup.Find("span", "class", "substitution-out").Pointer != nil {
						eventType = "substitution-out"
					}else if lineup.Find("span", "class", "y-card").Pointer != nil {
						eventType = "y-card"
					}else if lineup.Find("span", "class", "soccer-ball").Pointer != nil {
						eventType = "soccer-ball"
					}else if lineup.Find("span", "class", "substitution-in").Pointer != nil {
						eventType = "substitution-in"
					}else if lineup.Find("span", "class", "r-card").Pointer != nil {
						eventType = "r-card"
					}
					eventDescription = lineup.Attrs()["title"]
				}
				event := map[string]string{
					eventType: eventDescription,
				}

				events[currentTitleEvents] = append(events[currentTitleEvents], map[string]interface{}{
					"name": playerName,
					"href": playerHref,
					"event": event,
				})
			}

		} else {
			t := col.Find("td").Text()
			currentTitleEvents = t
		}
	}

	return json.MarshalIndent(events, " ", "	")

}