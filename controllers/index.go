package controllers

import "github.com/gin-gonic/gin"

func IndexIndex(c *gin.Context) {
	c.String(200, "roast server")
}