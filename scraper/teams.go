package scraper

import (
	"github.com/anaskhan96/soup"
	"strings"
	"log"
	"time"
	"strconv"
)

func GetFullTeamsInfo(tournamentUrl string, verbose bool) (map[string]*Team, time.Duration, error) {
	t0 := time.Now()

	resp, err := soup.Get(tournamentUrl)
	if err != nil {
		return map[string]*Team{}, time.Now().Sub(t0), err
	}
	doc := soup.HTMLParse(resp)

	participants := doc.Find("div", "id", "tournament-page-participants")
	body := participants.Find("table").Find("tbody")

	participantsData := make(map[string]*Team)
	teamName := ""

	currentTeam := new(Team)
	for _, p := range body.FindAll("tr") {
		pLink := p.Find("a")
		teamName = strings.ToLower(pLink.Text())
		if verbose {
			log.Println(teamName)
		}
		teamUrl := BaseScoreBoard + pLink.Attrs()["href"]
		plan := teamUrl + "/plantilla"

		r, err := soup.Get(plan)
		if err != nil {
			return map[string]*Team{}, time.Now().Sub(t0), err
		}
		planSoup := soup.HTMLParse(r)
		playerTable := planSoup.Find("table", "class", "base-table squad-table")
		field := ""



		for _, row := range  playerTable.Find("tbody").FindAll("tr") {

			jerseyNumber := 0
			name := ""
			nation := ""
			playerAge := 0
			played := ""
			goals := 0
			yellows := 0
			reds := 0


			innerClass := row.Attrs()["class"]
			if strings.Contains(innerClass, "player-type-title") {
				field = strings.ToLower(row.Find("td").Text())
			} else {
				for i, td := range row.FindAll("td") {
					if  td.Pointer.FirstChild != nil {
						switch i {
						case 0:
							jerseyNumberS := td.Text()
							jerseyNumber, _ = strconv.Atoi(jerseyNumberS)
							break
						case 1:
							name = td.Find("a").Text()
							nation = td.Find("span").Attrs()["title"]
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
				}
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
					currentTeam.MidFielder = append(currentTeam.MidFielder, p)
				case "defensas":
					currentTeam.Defender = append(currentTeam.Defender, p)
				case "delanteros":
					currentTeam.Forwarder = append(currentTeam.Forwarder, p)
				case "entrenador":
					currentTeam.Coach = p
				case "porteros":
					currentTeam.GolKeeper = append(currentTeam.GolKeeper, p)

				}

			}
		}
		participantsData[teamName] = currentTeam
	}

	t1 := time.Now()
	return participantsData, t1.Sub(t0), err
}