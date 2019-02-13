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
        v1.GET("brew-methods", BrewMethodsGetBrewMethods)
        v1.POST("cafes/:id/like", CafesPostLikeCafe)
        v1.DELETE("cafes/:id/like", CafesDeleteLikeCafe)
        v1.POST("cafes/:id/tags", CafesPostAddTags)
        v1.DELETE("cafes/:id/tags/:tagId", CafesDeleteCafeTag)
    }

	return router
}