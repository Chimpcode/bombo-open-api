package scraper

import (
	"log"
	"strings"

	"regexp"

	"github.com/PuerkitoBio/goquery"
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

	finalEvents := MatchEvents{Home: map[string][]EventPlayer{}, Away: map[string][]EventPlayer{}}

	c := colly.NewCollector()

	// Only lineups
	c.OnHTML("div.lineups-wrapper > table.parts > tbody", func(e *colly.HTMLElement) {
		typeOfEvent := ""

		lEventsPlayer := make([]EventPlayer, 0)
		rEventsPlayer := make([]EventPlayer, 0)

		firstTimeFlag := true
		e.DOM.Find("tr").Each(func(i int, s *goquery.Selection) {
			class, ifExist := s.Attr("class")
			if ifExist && (class == "odd" || class == "even") {
				colLeft := s.Find("td.fl")
				colRight := s.Find("td.fr")

				// LEFT
				lName := colLeft.Find("div.name").Find("a").Text()
				lNation, _ := colLeft.Find("span").Attr("title")
				lJersey := colLeft.Find("div.time-box").Text()

				lEvents := make([]Event, 0)

				colLeft.Find(".icon-lineup").Each(func(n int, s *goquery.Selection) {
					rawData, exist := s.Attr("title")
					if !exist {
						return
					}

					classes, ifTypeExist := s.Find("span").Attr("class")

					classes = strings.Replace(classes, " ", "", -1)
					typeClass := strings.Replace(classes, "icon", "", -1)

					if !ifTypeExist {
						return
					}

					time := strings.Split(rawData, "'")[0] + "'"

					lEvents = append(lEvents, Event{
						At:       time,
						Count:    n,
						Metadata: rawData,
						Type:     typeClass,
					})
				})

				lEventsPlayer = append(lEventsPlayer, EventPlayer{
					Name:   lName,
					Nation: lNation,
					Jersey: lJersey,
					Events: lEvents,
				})

				// RIGHT
				rName := colRight.Find("div.name").Find("a").Text()

				rNation, _ := colRight.Find("span").Attr("title")
				rJersey := colRight.Find("div.time-box").Text()

				rEvents := make([]Event, 0)

				colRight.Find(".icon-lineup").Each(func(n int, s *goquery.Selection) {
					rawData, exist := s.Attr("title")
					if !exist {
						return
					}

					extraReg, err := regexp.Compile(`(\().+(\))`)
					if err != nil {
						return
					}
					result := extraReg.Find([]byte(rawData))

					log.Println("====>", string(result))

					classes, ifTypeExist := s.Find("span").Attr("class")

					classes = strings.Replace(classes, " ", "", -1)
					typeClass := strings.Replace(classes, "icon", "", -1)

					if !ifTypeExist {
						return
					}

					time := strings.Split(rawData, "'")[0] + "'"

					rEvents = append(rEvents, Event{
						At:       time,
						Count:    n,
						Metadata: rawData,
						Type:     typeClass,
					})
				})

				rEventsPlayer = append(rEventsPlayer, EventPlayer{
					Name:   rName,
					Nation: rNation,
					Jersey: rJersey,
					Events: rEvents,
				})

				log.Println(lName, rName)
			} else {
				if !firstTimeFlag {
					finalEvents.Home[typeOfEvent] = lEventsPlayer
					finalEvents.Away[typeOfEvent] = rEventsPlayer

					lEventsPlayer = make([]EventPlayer, 0)
					rEventsPlayer = make([]EventPlayer, 0)

					firstTimeFlag = false
				} else {
					firstTimeFlag = false
				}
				typeOfEvent = s.Find("td").Text()

				typeOfEvent = ParseToValidCamelCase(typeOfEvent)

				log.Printf("======== %s ========", typeOfEvent)
			}
		})

		finalEvents.Home[typeOfEvent] = lEventsPlayer
		finalEvents.Away[typeOfEvent] = rEventsPlayer
	})

	err := c.Post(finalURL, map[string]string{})
	if err != nil {
		return finalEvents, err
	}

	return finalEvents, nil
}
