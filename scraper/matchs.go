package scraper

import (
	"log"
	"strings"

	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/k0kubun/pp"
	"net/http"
	"bytes"
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

	pp.Println(finalURL)

	finalEvents := MatchEvents{
		Home: BaseMatchEvent{Events: map[string][]EventPlayer{}},
		Away: BaseMatchEvent{Events: map[string][]EventPlayer{}},
	}
	c := colly.NewCollector()

	// Only lineups
	c.OnHTML("div.lineups-wrapper > table.parts > tbody", func(e *colly.HTMLElement) {
		pp.Println("INTO DIV.LINE...")
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
					finalEvents.Home.Events[typeOfEvent] = lEventsPlayer
					finalEvents.Away.Events[typeOfEvent] = rEventsPlayer

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

		finalEvents.Home.Events[typeOfEvent] = lEventsPlayer
		finalEvents.Away.Events[typeOfEvent] = rEventsPlayer
	})

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("x-fsign", "vPAvOgxk")
		pp.Println("(MATCH) Visiting", r.URL.String())
		pp.Println("(MATCH) Visiting [HEADERS]", r.Headers)
		pp.Println("(MATCH) Visiting [BODY]", r.Body)
	})

	c.OnResponse(func(r *colly.Response) {
		pp.Println(r.StatusCode)
		pp.Println(string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//err := c.Post(finalURL, map[string]string{})
	//if err != nil {
	//	log.Println("c.Post(finalURL, map[string]string{}) ", err)
	//	return finalEvents, err
	//}

	//c.Visit(finalURL)
	body := bytes.NewBufferString("")
	err := c.Request("GET", finalURL, body, nil, http.Header{"x-fsign":[]string{"SW9D1eZo"}})
	if err != nil {
		log.Println("c.Request('GET', finalURL ...) ", err)
		return finalEvents, err
	}

	score, err := GetScoreFromMatch(urlMatch)
	if err != nil {
		return finalEvents, err
	}
	finalEvents.Home.Name = score.Home.Name
	finalEvents.Away.Name = score.Away.Name

	finalEvents.Home.Score = score.Home.Score
	finalEvents.Away.Score = score.Away.Score

	pp.Println(finalEvents)

	return finalEvents, nil
}
