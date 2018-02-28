package main

import "time"

const WorkMatchType = "match"
const WorkLeagueType = "league"

const WorkWaiting = "waiting"
const WorkUpdating = "updating"

type Work struct {
	URL string `json:"url"`
	Name string `json:"name"`
	Period time.Duration `json:"period"`
	Type string `json:"type"`
	LastUpdate time.Time `json:"last_update"`
	State string `json:"state"`
	LastError error `json:"last_error"`
}
