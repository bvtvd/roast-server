

func(c *gin.Context) {
	 claims := c.MustGet("claims").(*jwtauth.CustomClaims)
	 fmt.Println(claims.ID)
	 c.String(http.StatusOK, claims.Name)
 }