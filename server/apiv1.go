package server

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"encoding/json"
)

func LinkApi(api iris.Party) {

	api.Get("/full", func(c iris.Context) {
		data, err := ioutil.ReadFile("./data/premier_league.json")
		if err != nil {
			c.JSON(iris.Map{
				"error_at": 0,
				"error": err.Error(),
			})
			return
		}
		var finalJson interface{}
		err = json.Unmarshal(data, &finalJson)
		if err != nil {
			c.JSON(iris.Map{
				"error_at": 1,
				"error": err.Error(),
			})
			return
		}
		_, err = c.JSON(finalJson)
		if err != nil {
			c.JSON(iris.Map{
				"error_at": 2,
				"error": err.Error(),
			})
			return
		}
		return
	})

}
