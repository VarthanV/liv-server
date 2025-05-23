package main

import (
	"github.com/VarthanV/liv-server/controllers"
	"github.com/gin-gonic/gin"
)

func initServer() {
	r := gin.Default()
	controller := controllers.NewController()
	controller.InitRoutes(r)
}

func main() {
	initServer()
}
