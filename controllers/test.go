package controllers

import (
	"github.com/gin-gonic/gin"
	"roast-server/utils"
)


// 获取所有咖啡店
func Geocode(c *gin.Context) {
	coordinate, _ := utils.GeocodeAddress("天城路1号", "杭州", "浙江")

	c.JSON(200, coordinate)
}