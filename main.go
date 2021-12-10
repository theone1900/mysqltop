package main

import (
	"github.com/gin-gonic/gin"
	"mysqltop/config"
	"mysqltop/modules/server"
	"mysqltop/routes"
)

func main() {
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//conn.GetDBSource(0,0,"")
	router := routes.InitRouter()

	server.Run(router)

}
