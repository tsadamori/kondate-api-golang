package main

import "github.com/tsadamori/kondate-api-golang/controller"

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}