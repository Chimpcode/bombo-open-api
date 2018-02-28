package scraper

import (
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
	"io/ioutil"
)

const costsFile = "./files/league_csv.csv"
var leagueCosts = map[string]float64{}

func init() {
	leagueCosts = make(map[string]float64)

	data, err := ioutil.ReadFile(costsFile)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		line = strings.Trim(line, "\r")
		chunks := strings.Split(line, ",")
		if len(chunks) > 2 {
			stringCost := chunks[2]
			cost, err := strconv.ParseFloat(stringCost, 64)
			if err != nil {
				continue
			}
			leagueCosts[chunks[0]] = cost
		}else{
			continue
		}

	}

	log.Println("Players costs loaded!")
}

func GetFullTeamsInfo(tournamentUrl string, verbose bool) (League, error) {

	log.Println("I'm inside")

	finalLeague := League{}

	c := colly.NewCollector()


	c.OnHTML("div#tournament-page-participants", func(e *colly.HTMLElement) {
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

		teamName := e.DOM.Find("div.team-name").First().Text()

		field := ""

		team := new(Team)

		columns :=e.DOM.Find("table.base-table").First().Find("tbody").Find("tr")


		columns.Each(func(i int, s *goquery.Selection) {
			if s.HasClass("player-type-title") {
				f := s.Find("td").First().Text()
				field = strings.ToLower(f)
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
						switch j {
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
					Cost: leagueCosts[name],

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

	return finalLeague[:len(finalLeague)-1], nil
}