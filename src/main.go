package main

import (
	"fmt"
	"im-stub/config"
	"im-stub/route"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	configs, _ := config.ReadConfig("../config/config.json")
	for index, routeInfo := range configs.RouteList {
		route.NewRoute(server, routeInfo)
		fmt.Println(index, routeInfo)
	}
	if !configs.EnableLog {
		gin.SetMode(gin.ReleaseMode)
	}
	serverRun := fmt.Sprintf("%s:%d", configs.Host, configs.Port)
	server.Run(serverRun)
}
