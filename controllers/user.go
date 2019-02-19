package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	// "roast-server/utils"
	// "roast-server/utils/jwt"
	"fmt"
	// "strconv"
	// validator "gopkg.in/go-playground/validator.v9"
	"net/http"
)

func UsersGetUser(c *gin.Context) {
	 // claims := c.MustGet("claims").(*jwt.CustomClaims)
	 user := c.MustGet("user").(User)
	 // fmt.Println(claims.ID)
	 fmt.Printf("user: %v", user)
	 c.JSON(http.StatusOK, user)
}