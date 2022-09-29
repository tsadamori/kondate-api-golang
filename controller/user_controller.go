package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/model"
)

func ShowAllUsers(c *gin.Context) {
	data := model.GetAllUsers()
	c.JSON(200, data)
}