package scraper

import (
	"github.com/anaskhan96/soup"
	"github.com/k0kubun/pp"
	"strings"
	"encoding/json"
)

func GetFullTeamsInfo() ([]byte, error) {
	tournamentUrl := TournamentA

	resp, err := soup.Get(tournamentUrl)
	if err != nil {
		return []byte{}, err
	}
	doc := soup.HTMLParse(resp)

	participants := doc.Find("div", "id", "tournament-page-participants")
	body := participants.Find("table").Find("tbody")

	participants_data := map[string]map[string][]map[string]string{}
	team_name := ""
	for _, p := range body.FindAll("tr") {
		p_link := p.Find("a")
		team_name = strings.ToLower(p_link.Text())
		pp.Println(team_name)
		team_url := BaseScoreBoard + p_link.Attrs()["href"]
		plan := team_url + "/plantilla"

		r, err := soup.Get(plan)
		if err != nil {
			return []byte{}, err
		}
		plan_soup := soup.HTMLParse(r)
		player_table := plan_soup.Find("table", "class", "base-table squad-table")
		field := ""
		players := map[string][]map[string]string{}

		for _, row := range  player_table.Find("tbody").FindAll("tr") {

			jersey_number := ""
			name := ""
			nation := ""
			player_age := ""
			played := ""
			goals := ""
			yellows := ""
			reds := ""


			innerClass := row.Attrs()["class"]
			if strings.Contains(innerClass, "player-type-title") {
				field = strings.ToLower(row.Find("td").Text())
				players[field]= []map[string]string{}
			} else {
				for i, td := range row.FindAll("td") {
					if  td.Pointer.FirstChild != nil {
						switch i {
						case 0:
							jersey_number = td.Text()
							break
						case 1:
							name = td.Find("a").Text()
							nation = td.Find("span").Attrs()["title"]
							break
						case 2:
							player_age = td.Text()
							break
						case 3:
							played = td.Text()
							break
						case 4:
							goals = td.Text()
							break
						case 5:
							yellows = td.Text()
							break
						case 6:
							reds = td.Text()
							break
						default:
							break
						}
					}
				}
				meta_data := map[string]string{
					"j_number": jersey_number,
					"name": name,
					"nation": nation,
					"player_age": player_age,
					"played": played,
					"goals": goals,
					"yellows": yellows,
					"reds": reds,
				}
				players[field] = append(players[field], meta_data)
			}
		}
		participants_data[team_name] = players
	}

	data, err := json.Marshal(participants_data)

	return data, err
}