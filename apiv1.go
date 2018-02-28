package main

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"encoding/json"
)

func LinkApi(api iris.Party, manager *WorkManager) {

	api.Get("/{endpoint: string}", func(c iris.Context) {
		endpoint := c.Params().Get("endpoint")
		for _, work := range manager.Works {
			if work.Name == endpoint {
				filePath := "./data/" + work.Type + "_" +work.Name + ".json"

				data, err := ioutil.ReadFile(filePath)
				if err != nil{
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

	})

}