package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	"fmt"
)



// 获取所有咖啡店
func CafesGetCafes(c *gin.Context) {

	cafes := GetAllCafes()
	fmt.Printf("cafes %v", cafes)

	// c.JSON(200, cafes)
	c.String(200, "hello")
}


func CafesPostNewCafe(c *gin.Context) {

}

func CafesGetCafe(c *gin.Context) {
	// id := c.Param("id")


}