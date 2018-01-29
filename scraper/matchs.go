package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func GetScoreFromMatch(urlMatch string) (MatchScore, error) {
	//
	scoreHome := ""
	scoreAway := ""
	teamHome := ""
	teamAway := ""

	c := colly.NewCollector()

	c.OnHTML("td#flashscore_column", func(e *colly.HTMLElement) {
		teamHome = e.DOM.Find("td.tname-home").Find("a").Text()
		teamAway = e.DOM.Find("td.tname-away").Find("a").Text()
	})

	c.OnHTML("td.current-result", func(e *colly.HTMLElement) {
		spans := e.DOM.Find("span")
		if spans.Length() != 3 {
			return
		}
		score := spans.Text()

		sScore := strings.Split(score, ":")
		scoreHome = sScore[0]
		scoreAway = sScore[len(sScore)-1]

	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
		return
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting ", r.URL)
	})

	c.Visit(urlMatch)

	match := MatchScore{
		Home: TeamBase{Name: teamHome, Score: scoreHome},
		Away: TeamBase{Name: teamAway, Score: scoreAway},
	}

	return match, nil

}

// GetEventsFromMatch ...
func GetEventsFromMatch(urlMatch string) (MatchEvents, error) {
	matchID := GetIdFromMatchUrl(urlMatch)
	finalURL := InternalScoreBoardApi + "d_li_" + matchID + "_es_1"

	finalEvents := new(MatchEvents)

	c := colly.NewCollector()

	// Only lineups
	c.OnHTML("div.lineups-wrapper > table.parts > tbody", func(e *colly.HTMLElement) {
		log.Println(e.DOM.Find("tr").Length())

	})

	err := c.Post(finalURL, map[string]string{})
	if err != nil {
		return *finalEvents, err
	}

	return *finalEvents, nil
}
