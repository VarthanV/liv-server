package controllers

import (
	"fmt"

	"github.com/VarthanV/liv-server/fileservice"
	"github.com/gin-gonic/gin"
)

const (
	host           = "127.0.0.1"
	fileServerPort = 8060
	websocketPort  = 8070
)

type Controller struct {
	fileService *fileservice.Service
}

func NewController() *Controller {
	return &Controller{
		fileService: fileservice.NewFileService(),
	}
}

func (c *Controller) InitRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.GET("/*path", c.ServeFile)
	r.Run(fmt.Sprintf("%s:%d", host, fileServerPort))

}

func (c *Controller) InitSocket(r *gin.Engine) {
	r.GET("/ws", c.HandleSocket)
	r.Run(fmt.Sprintf("%s:%d", host, websocketPort))
}
