package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	go startApiServer()
	router := gin.Default()
	router.Static("/", "portal")
	router.Run(":1024")
}

func startApiServer() {
	apiServer := gin.Default()
	apiServer.GET("/qn/api/rtmp/create", func(context *gin.Context) {
		result := &Result{Message: "success", Code: 200, Result: "null"}
		context.JSON(200, result)
	})
	fmt.Println("starting api server .....")
	apiServer.Run(":1023")
}

type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Result  string `json:"result"`
}
