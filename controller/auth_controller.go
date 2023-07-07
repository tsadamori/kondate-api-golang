package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/auth"
	"github.com/tsadamori/kondate-api-golang/model"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	isValidUser := model.IsValidUser(email, password)
	if isValidUser == false {
		errorMessage := gin.H{
			"error": "Login failed.",
		}
		c.JSON(401, errorMessage)
	}

	token, err := auth.CreateJWTSignedToken(email)
	if err != nil {
		errorMessage := gin.H{
			"error": "Create JWT signed token failed.",
		}
		c.JSON(500, errorMessage)
	}
	c.SetCookie(auth.TokenKey, token, 3600, "/", "localhost", false, false)
}

func Logout(c *gin.Context) {
	c.SetCookie(auth.TokenKey, "", -3600, "/", "localhost", false, false)
}