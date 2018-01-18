from bs4 import BeautifulSoup
import requests
import json

base_uri = 'https://www.scoreboard.com'

tournamentA = 'https://www.scoreboard.com/es/futbol/inglaterra/championship/equipos'
tournamentB = 'https://www.scoreboard.com/es/futbol/inglaterra/premier-league/equipos'

r = requests.get(tournamentA)
html_data = r.text

soup = BeautifulSoup(html_data)
participants = soup.find(id='tournament-page-participants')
body = participants.table.find('tbody')

participants_data = {}

for p in body.find_all('tr'):
    p_link = p.find('a')
    team_name = p_link.get_text()
    team_url = base_uri + p_link.get('href')
    plan = team_url + '/plantilla'

    r = requests.get(plan)
    plan_soup = BeautifulSoup(r.text)
    player_table = plan_soup.find('table', {'class': 'base-table squad-table'})

    field = ''
    players = {}
    for row in player_table.find('tbody').find_all('tr'):

        if 'player-type-title' in row.get('class'):
            field = row.find('td').get_text().lower()
            players[field]= []
        else:
            for i, td in enumerate(row.find_all('td')):
                if i == 0:
                    jersey_number = td.get_text()
                elif i == 1:
                    name = td.find('a').get_text()
                    nation = td.find('span').get('title')
                elif i == 2:
                    player_age = td.get_text()
                elif i == 3:
                    played = td.get_text()
                elif i == 4:
                    goals = td.get_text()
                elif i == 5:
                    yellows = td.get_text()
                elif i == 6:
                    reds = td.get_text()
            meta_data = {
                'j_number': jersey_number,
                'name': name,
                'nation': nation,
                'player_age': player_age,
                'played': played,
                'goals': goals,
                'yellows': yellows,
                'reds': reds
            }
            players[field].append(meta_data)
    participants_data[team_name.lower()] =players

final_data = json.dumps(participants_data,indent=2, sort_keys=True)
with open('data.json', 'w') as handle:
    handle.write(final_data)

