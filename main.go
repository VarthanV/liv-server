package main

import (
	"github.com/VarthanV/liv-server/controllers"
	"github.com/gin-gonic/gin"
)

func initServer() {
	r := gin.Default()
	socketRouter := gin.Default()
	controller := controllers.NewController()
	go controller.InitSocket(socketRouter)
	controller.InitRoutes(r)
}

func main() {
	initServer()
}
