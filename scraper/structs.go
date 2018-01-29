package scraper

type TeamBase struct {
	Name  string
	Score string
}

type MatchScore struct {
	Home TeamBase
	Away TeamBase
}

type Event struct {
	Type     string
	Count    int
	Metadata string
	At       string
	Extras   map[string]string
}

type EventPlayer struct {
	Name       string
	Jersey     string
	NationFlag string

	Events []Event
}

type MatchEvents struct {
	Home map[string]EventPlayer
	Away map[string]EventPlayer
}
