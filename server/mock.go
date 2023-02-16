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
		fmt.Println("🎭 Start Mock Server")
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
	parsers, err := parser.FromYaml(config.Cfg.GetMockfilePath())
	if err != nil {
		panic(err)
	}
	// init all request handlers from parsers
	for _, p := range parsers {
		router.Handle(p.GetMethodType(), p.GetURL(),
			mock.GenHandler(p.GetMethodType(), p.GetResp(), p.GetStatus()))
	}
	fmt.Printf("🍺 MockServer is running on port :%d\n", config.Cfg.GetServerPort())
	err = router.Run(fmt.Sprintf(":%d", config.Cfg.GetServerPort()))
	if err != nil {
		panic(err)
	}
}

func runAsGen() {

}
