package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vanga-top/skyline-spider/api_server/config"
	"github.com/vanga-top/skyline-spider/pili"
)

/**
  --创建rtmp推流地址
	这边可以调用pili包下的service
*/
func CreateRTMPPushStreamURL() (string, gin.HandlerFunc) {
	fun := func(context *gin.Context) {
		err := pili.CreateRTMPPushURL(config.ApiServerConfig)
		if err != nil {
			return
		}
		result := &config.Result{Message: "success", Code: 200, Result: "null"}
		context.JSON(200, result)
	}
	return "/qn/api/rtmp/create", fun
}
