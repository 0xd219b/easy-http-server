package server

import (
	"fmt"

	"github.com/Berger7/easy-http-server/config"
	"github.com/Berger7/easy-http-server/pkg/mock"
	"github.com/Berger7/easy-http-server/pkg/parser"
	"github.com/gin-gonic/gin"
)

func RunMockServer() {
	if config.Cfg.GetMockService() == "server" {
		fmt.Println(config.Cfg.GetMockService())
		runAsServer()
	} else {
		fmt.Println(config.Cfg.GetMockService())
		runAsGen()
	}
}

func runAsServer() {
	// run server by gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	parsers, err := parser.ParserFromYaml(config.Cfg.GetMockfilePath())
	if err != nil {
		panic(err)
	}
	// init all request handlers from parsers
	for _, parser := range parsers {
		router.Handle(parser.GetMethodType(), parser.GetURL(),
			mock.GenHandler(parser.GetURL(), parser.GetMethodType(), parser.GetResp()))
	}
	err = router.Run(fmt.Sprintf(":%d", config.Cfg.GetServerPort()))
	if err != nil {
		panic(err)
	}
}

func runAsGen() {

}
