package  main

import (
	"github.com/kataras/iris"
	"errors"
	"time"
	"log"
)

func LinkManagerApi(managerApi iris.Party,  manager *WorkManager) {

	managerApi.Get("/get-works", func(c iris.Context) {
		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"data": manager.Works,
			"error": nil,
		})
	})

	managerApi.Post("/add-work", func(c iris.Context) {
		work := new(Work)

		err := c.ReadJSON(work)
		if err !=nil {
			log.Println(err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"data": nil,
				"error": err.Error(),
			})
			return
		}


		if work.Type != "" {
			if work.Type == WorkMatchType || work.Type == WorkLeagueType  {

				work.LastError = nil
				work.LastUpdate = time.Now()

				manager.AddWork(work)

				err := SaveWorkManagerState(manager)
				if err !=nil {
					log.Println(err)
					c.StatusCode(iris.StatusInternalServerError)
					c.JSON(iris.Map{
						"data": nil,
						"error": err.Error(),
					})
					return
				}
				c.StatusCode(iris.StatusOK)
				c.JSON(iris.Map{
					"data": manager.Works,
					"error": nil,
				})
				return
			}
		}

		c.StatusCode(iris.StatusInternalServerError)
		log.Println(errors.New("invalid work type"))
		c.JSON(iris.Map{
			"data": nil,
			"error": errors.New("invalid work type").Error(),
		})


	})

}
