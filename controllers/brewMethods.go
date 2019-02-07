package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	// "roast-server/utils"
	// "fmt"
	// "strconv"
	// validator "gopkg.in/go-playground/validator.v9"
)


// 获取所有冲泡方法
func BrewMethodsGetBrewMethods(c *gin.Context) {

	cafes := GetBrewMethods()
	// fmt.Printf("cafes %v", cafes)

	c.JSON(200, cafes)
	// c.String(200, "hello")
}