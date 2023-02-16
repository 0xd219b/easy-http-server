package mock

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenHandler(method string, resp []byte, status int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if status == 0 {
			switch method {
			case "GET":
				status = http.StatusOK
			case "POST":
				status = http.StatusCreated
			default:
				status = http.StatusOK
			}
		}
		var f interface{}
		err := json.Unmarshal(resp, &f)
		if err != nil {
			panic(err)
		}
		c.JSON(status, f)
	}
}
