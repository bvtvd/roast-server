package routers

import (
	"github.com/gin-gonic/gin"
	// . "roast-server/middlewares"
	"github.com/gin-contrib/cors"
)



func InitRouter() *gin.Engine {

	router := gin.Default()

	/**
	* 加载全局中间件
	*/
	// router.Use(Cors())	// 跨域
	router.Use(cors.Default())

	/**
	* 加载路由	
	*/
	NoVersion(router)
	V1(router)

	return router
}