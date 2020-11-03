package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/vanga-top/skyline-spider/api_server"
)

func main() {
	router := gin.Default()
	router.Static("/", "portal")
	router.Run(":1024")
}
