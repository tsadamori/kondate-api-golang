package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", ShowAllUsers)
	r.GET("/users/:id", ShowUserById)
	return r
}