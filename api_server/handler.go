package api_server

import "github.com/gin-gonic/gin"


func createRTMPURL() gin.HandlerFunc {
	fun := func(context *gin.Context) {
		result := &Result{Message: "success", Code: 200, Result: "null"}
		context.JSON(200, result)
	}
	return fun
}
