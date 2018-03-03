package main

import (
	"time"
	"log"
	//"github.com/k0kubun/pp"
	"io/ioutil"
	"encoding/json"
	"os"
)

type WorkManager struct {
	Works []*Work `json:"works"`
	WorksUpdating []*Work `json:"works_updating"`

}


func LoadManagerFromFile (filePath string) (*WorkManager, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &WorkManager{}, err
	}

	wm := new(WorkManager)

	err = json.Unmarshal(data, wm)
	if err != nil {
		return &WorkManager{}, err
	}
	return wm, nil
}

func SaveWorkManagerState(wm *WorkManager) error {
	data, err := json.MarshalIndent(wm, "", "	")
	if err != nil {
		return err
	}

	file, err := os.Create("./state/manager_state.json")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewWorkManager() *WorkManager {
	works := make([]*Work, 0)
	worksUpdating := make([]*Work, 0)

	return &WorkManager{
		Works: works,
		WorksUpdating: worksUpdating,
	}
}


func (manager *WorkManager) Run(step time.Duration) {
	go func() {
		for {
			time.Sleep(step)

			for i, work := range manager.WorksUpdating {
				log.Printf("Deleting (%d)", i)
				if work.State == WorkWaiting {
					if i<len(manager.WorksUpdating) {
						manager.WorksUpdating = append(
							manager.WorksUpdating[:i],
							manager.WorksUpdating[i+1:]...
						)
					}

				}
			}

			for _, work := range manager.Works {

				timeForUpdate := work.LastUpdate.Add(work.Period)

				if  timeForUpdate.Before(time.Now())  && work.State != WorkUpdating {
					log.Println("Updating ", work.Name)
					manager.WorksUpdating = append(manager.WorksUpdating, work)
					work.Update()

				}
			}
			//pp.Println(manager.WorksUpdating)

		}
	}()

}


func (manager *WorkManager) AddWork(work *Work) {
	manager.Works = append(manager.Works, work)
}

func (manager *WorkManager) DeleteWork(i int) {
	manager.Works = append(
		manager.Works[:i],
		manager.Works[i+1:]...
	)
}