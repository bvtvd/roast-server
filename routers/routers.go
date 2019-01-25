package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {

	router := gin.Default()

	/**
	* 加载路由	
	*/
	NoVersion(router)
	V1(router)

	return router
}