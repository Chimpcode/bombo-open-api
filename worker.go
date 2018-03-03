package main

import (
	"time"
	"github.com/k0kubun/pp"
)

func NewWork(url string, name string, typeOfWork string, period time.Duration) *Work {
	return &Work{
		URL: url,
		Name: name,
		Type: typeOfWork,
		Period: period,
		LastUpdate: time.Now(),
		LastError: nil,
		State: "created",
	}

}

func (w *Work) Update() {
	go func() {
		w.LastUpdate = time.Now()
		w.State = WorkUpdating

		switch w.Type {
		case WorkMatchType:
			pp.Println("UPDATING MATCH!!!")
			err := SaveDataForMatchWork(w)
			if err != nil {
				pp.Println(err)
				w.LastError = err
				break
			}
			w.LastError = nil
			break
		case WorkLeagueType:
			err := SaveDataForLeagueWork(w)
			if err != nil {
				w.LastError = err
				break
			}
			w.LastError = nil
			break
		}

		w.State = WorkWaiting

	}()
}