package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/auth"
	"github.com/tsadamori/kondate-api-golang/model"
)

type ErrorMessage struct {
	error string
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	isValidUser := model.IsValidUser(email, password)
	if isValidUser == false {
		errorMessage := ErrorMessage{
			error: "Login failed.",
		}
		c.JSON(401, errorMessage)
	}

	token, err := auth.CreateJWTSignedToken(email, password)
	if err != nil {
		errorMessage := ErrorMessage{
			error: "Create JWT signed token failed.",
		}
		c.JSON(500, errorMessage)
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, false)
}
