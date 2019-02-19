package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	// "roast-server/utils"
	"fmt"
	middlewares "roast-server/utils/jwt"
	// "strconv"
	// validator "gopkg.in/go-playground/validator.v9"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	// "time"
)

type Credentials struct {
	Username string	 `form:"username" json:"username"`
	Password string  `form:"password" json:"password"`
}

func AuthLogin(c *gin.Context) {
	var credentials Credentials
	c.ShouldBindJSON(&credentials)

	fmt.Printf("credentials: %v\n", credentials)
	user, err := GetUserFromCredentials(User{Username: credentials.Username, Password: credentials.Password})
	if err != nil {
		c.JSON(400, gin.H{"msg": "用户名或密码错误"})
	} else {

		// 用户验证成功, 准备token 然后返回
		j := &middlewares.JWT{
			[]byte("secret"),
		}

		claims := middlewares.CustomClaims{
			int(user.ID),
			user.Name,
			jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer:    "roast",
			},
		}

		token, err := j.CreateToken(claims)
		fmt.Printf("token: %s\n", token)
		if err != nil {
			c.JSON(400, err.Error())
		}
		
		c.JSON(http.StatusOK, gin.H{"token": token})
	}

}
