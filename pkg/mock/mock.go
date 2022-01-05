package mock

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenHandler(url, method string, resp string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if method == "GET" {
			var f interface{}
			err := json.Unmarshal([]byte(resp), &f)
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, f)
		} else if method == "POST" {
			var f interface{}
			err := json.Unmarshal([]byte(resp), &f)
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusCreated, f)
		}
	}
}
