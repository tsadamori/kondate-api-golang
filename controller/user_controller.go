package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/model"
)

func ShowAllUsers(c *gin.Context) {
	datas := model.GetAllUsers()
	fmt.Println(datas)
}