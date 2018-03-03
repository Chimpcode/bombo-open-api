package scraper

type TeamBase struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

type MatchScore struct {
	Home TeamBase `json:"home"`
	Away TeamBase `json:"away"`
}

type Event struct {
	Type     string `json:"type"`
	Count    int `json:"count"`
	Metadata string `json:"metadata"`
	At       string `json:"at"`
	Extras   map[string]string `json:"extras"`
}

type EventPlayer struct {
	Name   string `json:"name"`
	Jersey string `json:"jersey"`
	Nation string `json:"nation"`

	Events []Event `json:"events"`
}

type BaseMatchEvent struct {
	Name string `json:"name"`
	Score string `json:"score"`
	Events map[string][]EventPlayer `json:"events"`
}

type MatchEvents struct {
	Home BaseMatchEvent `json:"home"`

	Away BaseMatchEvent `json:"away"`
}
