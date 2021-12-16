package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Berger7/easy-http-server/config"
	"github.com/gin-gonic/gin"
)

const (
	defaultResponseBodyString = "hi, this resp from easy-http-server"
)

// RunSingleServer use gin to run a single server
func RunSingleServer() {
	switch config.Cfg.RequestMethod {
	case "GET":
		singleGetServer(config.Cfg.GetRequestURL(), defaultResponseBodyString)
	case "POST":
		singlePostServer(config.Cfg.GetRequestURL(), defaultResponseBodyString)
	default:
		fmt.Println("request method is not supported")
	}
}

func singleGetServer(url, body string) {
	// run server by gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET(url, func(c *gin.Context) {
		c.String(http.StatusOK, body)
	})
	router.Run(fmt.Sprintf(":%d", config.Cfg.ServerPort))
}

func singlePostServer(url, body string) {
	// run server by gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST(url, func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, "")
		} else {
			fmt.Println(string(b))
			c.String(http.StatusOK, body)
		}
	})
	router.Run(fmt.Sprintf(":%d", config.Cfg.ServerPort))
}
