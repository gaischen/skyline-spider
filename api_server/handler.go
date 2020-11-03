package api_server

import "github.com/gin-gonic/gin"

func createRTMPURL() gin.HandlerFunc {
	fun := func(context *gin.Context) {
		token := context.Query("token")
		if token == "" || token != ApiServerConfig.Token {
			result := &Result{Message: "unauth", Code: 403, Result: "null"}
			context.JSON(403, result)
			return
		}
		result := &Result{Message: "success", Code: 200, Result: "null"}
		context.JSON(200, result)
	}
	return fun
}

func RegisterHandler(uri string, handler gin.HandlerFunc) {

}

func test() {
	RegisterHandler("test", createRTMPURL())
}
