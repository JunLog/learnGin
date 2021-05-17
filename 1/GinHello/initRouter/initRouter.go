package initRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context){
		context.String(http.StatusOK, "hello gin")
	})
	return router
}