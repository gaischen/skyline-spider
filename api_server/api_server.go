package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//load config form file system
func init() {
	parserConfig()
	go startApiServer()
}

func startApiServer() {
	apiServer := gin.Default()
	apiServer.GET("/qn/api/rtmp/create", createRTMPURL())
	fmt.Println("starting api server successful.....")
	apiServer.Run(":1023")
}
