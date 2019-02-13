package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	// "roast-server/utils"
	// "fmt"
	// "strconv"
	// validator "gopkg.in/go-playground/validator.v9"
)


// 搜索标签
func TagsGetTags(c *gin.Context) {
	name := c.DefaultQuery("name", "")

	tags := GetTags(name)

	c.JSON(200, tags)
}