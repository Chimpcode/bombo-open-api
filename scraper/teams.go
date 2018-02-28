package scraper

import (
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

func GetFullTeamsInfo(tournamentUrl string, verbose bool) (*League, error) {

	log.Println("I'm inside")

	finalLeague := League{}

	c := colly.NewCollector()


	c.OnHTML("div#tournament-page-participants", func(e *colly.HTMLElement) {
		log.Println("I'm here!")
		e.DOM.Find("tr").Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").First().Attr("href")
			if exist {
				teamP := BaseScoreBoard + href + "/plantilla"
				log.Println("Visting: ", teamP)
				c.Visit(teamP)
			}
		})
	})

	c.OnHTML("div.main", func(e *colly.HTMLElement) {
		log.Println("I'm here! x22")

		teamName := e.DOM.Find("div.team-name").First().Text()
		log.Println(teamName)

		field := ""

		team := new(Team)

		e.DOM.Find("div#block-summary-squad >tbody>tr").Each(func(i int, s *goquery.Selection) {
			if s.HasClass("player-type-title") {
				log.Println(field)
				field, _ = s.Attr("class")
			}else{
				jerseyNumber := 0
				name := ""
				nation := ""
				playerAge := 0
				played := ""
				goals := 0
				yellows := 0
				reds := 0
				s.Find("td").Each(func(j int, td *goquery.Selection) {
					if  td.First() != nil {
						switch i {
						case 0:
							jerseyNumberS := td.Text()
							jerseyNumber, _ = strconv.Atoi(jerseyNumberS)
							break
						case 1:
							name = td.Find("a").Text()
							nation, _ = td.Find("span").Attr("title")
							break
						case 2:
							playerAgeS := td.Text()
							playerAge, _ = strconv.Atoi(playerAgeS)
							break
						case 3:
							played = td.Text()
							break
						case 4:
							goalsS := td.Text()
							goals, _ = strconv.Atoi(goalsS)
							break
						case 5:
							yellowsS := td.Text()
							yellows, _ = strconv.Atoi(yellowsS)
							break
						case 6:
							redsS := td.Text()
							reds, _ = strconv.Atoi(redsS)
							break
						default:
							break
						}
					}
				})

				nCode := ""
				nationCodes := strings.Split(NationsCode[nation], "_")
				if len(nationCodes) > 1 {
					nCode = strings.ToLower(nationCodes[1])
				} else {
					nCode = ""
				}

				p := &Player{
					Name: name,
					JNumber: jerseyNumber,
					Nation: nation,
					NationCode: nCode,
					Age: playerAge,
					Played: played,
					Goals: goals,
					Yellows: yellows,
					Reds: reds,
					Team: teamName,
					Cost: 0,

				}

				switch field {
				case "centrocampistas":
					team.MidFielder = append(team.MidFielder, p)
				case "defensas":
					team.Defender = append(team.Defender, p)
				case "delanteros":
					team.Forwarder = append(team.Forwarder, p)
				case "entrenador":
					team.Coach = p
				case "porteros":
					team.GolKeeper = append(team.GolKeeper, p)

				}

			}
		})
		team.Name = teamName

		finalLeague = append(finalLeague, team)
	})

	c.Visit(tournamentUrl)

	return &finalLeague, nil
}