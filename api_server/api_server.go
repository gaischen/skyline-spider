package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vanga-top/skyline-spider/api_server/config"
	"github.com/vanga-top/skyline-spider/api_server/controller"
)

//load config form file system
func init() {
	//解析config
	config.ParserConfig()
	register()
	go startApiServer()
}

/**
--注册controller,把逻辑和一些通用行为隔离开
*/
func register() {
	config.RegisterHandler(controller.CreateRTMPPushStreamURL())
}

func startApiServer() {
	apiServer := gin.Default()
	config.LoadHandler(apiServer)
	fmt.Println("starting api server successful.....")
	apiServer.Run(":1023")
}
