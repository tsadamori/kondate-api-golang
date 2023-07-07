package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsadamori/kondate-api-golang/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	namespace := "/api/v1"

	r.POST(namespace + "/login", Login)
	r.Use(middleware.AuthMiddleware)
	{
		r.GET(namespace + "/users", ShowAllUsers)
		r.GET(namespace + "/users/:id", ShowUserById)
		r.POST(namespace + "/users", RegisterUser)
	}
	return r
}