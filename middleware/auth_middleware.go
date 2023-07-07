package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/auth"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie(auth.TokenKey)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	_, err = auth.ParseJWTToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid Token",
		})
		c.Abort()
		return
	}

	c.Next()
}