package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	// "roast-server/utils"
	middlewares "roast-server/utils/jwt"
	"fmt"
	// "strconv"
	// validator "gopkg.in/go-playground/validator.v9"
	jwt "github.com/dgrijalva/jwt-go"
)

func AuthLogin(c *gin.Context) {
	username := c.PostForm("username");
	password := c.PostForm("password");

	fmt.Printf("username: %s\n", username)
	fmt.Printf("password: %s\n", password)
	user, err := GetUserFromCredentials(User{Username: username, Password: password})
	if err != nil {
		c.JSON(200, gin.H{"msg": "用户名或密码错误"})
	}else{
		// 用户验证成功, 准备token 然后返回
		j := &middlewares.JWT{
			[]byte("hello"),
		}

		claims := middlewares.CustomClaims{
			int(user.ID),
			user.Name,
			jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer: "roast",
			},
		}

		token, err := j.CreateToken(claims)
		if err != nil {
			c.JSON(204, err.Error())
			c.Abort()
		}else{
			c.JSON(201, gin.H{"token": token})
		}
	}
	
}