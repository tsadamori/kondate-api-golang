package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/model"
)

func ShowAllUsers(c *gin.Context) {
	var limit, offset int

	limit, _ = strconv.Atoi(c.Query("limit"))
	offset, _ = strconv.Atoi(c.Query("offset"))
	if limit == 0 {
		limit = 10
	}
	if offset == 0 {
		offset = 0
	}

	data := model.GetAllUsers(limit, offset)
	c.JSON(200, data)
}

func ShowUserById(c *gin.Context) {
	id := c.Param("id")
	data := model.GetUserById(id)
	c.JSON(200, data)
}

