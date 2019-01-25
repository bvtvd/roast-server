package routers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/controllers"
)

func NoVersion(router *gin.Engine) *gin.Engine {

	router.GET("/", IndexIndex)

	return router
}