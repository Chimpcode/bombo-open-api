package main

type Player struct {
	JNumber    int    `json:"j_number"`
	Name       string `json:"name"`
	Nation     string `json:"nation"`
	NationCode string `json:"nation_code"`
	Age        int    `json:"age"`
	Played     string `json:"played"`
	Goals      int    `json:"goals"`
	Yellows    int    `json:"yellows"`
	Reds       int    `json:"reds"`
	Team       string `json:"team"`
	Cost       int    `json:"cost"`
}

type Coach struct {
	Name       string `json:"name"`
	Nation     string `json:"nation"`
	NationCode string `json:"nation_code"`
	Age        int    `json:"age"`
	Team       string `json:"team"`
}

type Team struct {
	Name       string    `json:"name"`
	GolKeeper  []*Player `json:"gol_keeper"`
	MidFielder []*Player `json:"mid_fielder"`
	Defender   []*Player `json:"defender"`
	Forwarder  []*Player `json:"forwarder"`
	Coach      *Player   `json:"coach"`
}
