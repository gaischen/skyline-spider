package config

import (
	"github.com/gin-gonic/gin"
)

var handlerMap = make(map[string]gin.HandlerFunc)

func RegisterHandler(uri string, handler gin.HandlerFunc) { //处理一些通用逻辑
	fun := func(ctx *gin.Context) {
		token := ctx.Query("token")
		if token == "" || token != ApiServerConfig.Token {
			result := &Result{Message: "unauth", Code: 403, Result: "null"}
			ctx.JSON(403, result)
			return
		}
		handler(ctx)
	}
	//register
	handlerMap[uri] = fun
}

func LoadHandler(engine *gin.Engine) {
	for k, v := range handlerMap {
		engine.GET(k, v)
	}
}
