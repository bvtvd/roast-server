package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	// "fmt"
	"strconv"
)



// 获取所有咖啡店
func CafesGetCafes(c *gin.Context) {

	cafes := GetAllCafes()
	// fmt.Printf("cafes %v", cafes)

	c.JSON(200, cafes)
	// c.String(200, "hello")
}

// 创建新的咖啡店
func CafesPostNewCafe(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	city := c.PostForm("city")
	state := c.PostForm("state")
	zip := c.PostForm("zip")

	cafe := Cafe{Name: name, Address: address, City: city, State: state, Zip: zip}

	PostNewCafe(&cafe)

	c.JSON(201, cafe)
}

// 获取咖啡店详情
func CafesGetCafe(c *gin.Context) {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	cafe := GetCafe(idInt)

	c.JSON(200, cafe)
}

