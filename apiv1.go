package main

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"encoding/json"
	"github.com/k0kubun/pp"
	"log"
	"errors"
)

func LinkApi(api iris.Party, manager *WorkManager) {
	api.Get("/get-endpoints", func(c iris.Context) {
		type MiniWork struct {
			Name string `json:"name"`
			URL string `json:"url"`
			Type string `json:"type"`
			Endpoint string `json:"endpoint"`
		}

		miniWorks := make([]MiniWork, 0)

		for _, w := range manager.Works {
			if w.Type == "match" {
				mw := MiniWork{
					Name: w.Name,
					URL: w.URL,
					Type: w.Type,
					Endpoint: "/api/v1.0/" + w.Name,
				}
				miniWorks = append(miniWorks, mw)
			}
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"data": miniWorks,
			"error": nil,
		})

	})
	api.Get("/{endpoint: string}", func(c iris.Context) {
		endpoint := c.Params().Get("endpoint")
		pp.Println(endpoint)
		for _, work := range manager.Works {
			if work.Name == endpoint {
				filePath := "./data/" + work.Type + "_" +work.Name + ".json"

				pp.Println(filePath)
				data, err := ioutil.ReadFile(filePath)
				if err != nil{
					log.Println("Error at ioutil.ReadFile()")
					c.StatusCode(iris.StatusInternalServerError)
					c.JSON(iris.Map{
						"data": nil,
						"error": err,
					})
					return
				}

				var response interface{}

				err = json.Unmarshal(data, &response)
				if err != nil{
					c.StatusCode(iris.StatusInternalServerError)
					c.JSON(iris.Map{
						"data": nil,
						"error": err,
					})
					return
				}

				c.JSON(iris.Map{
					"data": response,
					"error": nil,
				})
				return
			}
		}

		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(iris.Map{
			"data": nil,
			"error": errors.New("work not found, check uri endpoint").Error(),
		})
		return


	})

}
