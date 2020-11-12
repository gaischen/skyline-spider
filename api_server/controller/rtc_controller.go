package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vanga-top/skyline-spider/api_server/config"
	"github.com/vanga-top/skyline-spider/pili"
)

func GetRTCRoomToken() (string, gin.HandlerFunc) {
	fun := func(context *gin.Context) {
		token, err := pili.GetRoomToken(config.ApiServerConfig)
		if err != nil {
			result := &config.Result{Message: "failed to get roomToken", Code: -1, Result: ""}
			context.JSON(-1, result)
			return
		}
		result := &config.Result{Message: "success", Code: 200, Result: token}
		context.JSON(200, result)
	}
	return "/qn/api/rtc/room/token", fun
}
