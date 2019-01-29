package routers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/controllers"
)

func V1(router *gin.Engine) *gin.Engine {

	// 简单组： v1
    v1 := router.Group("/v1")
    {
        v1.GET("/cafes", CafesGetCafes)
        v1.POST("/cafes", CafesPostNewCafe)
        v1.GET("/cafes/:id", CafesGetCafe)
        v1.GET("geocode", Geocode)
    }

	return router
}