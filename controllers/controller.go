package controllers

import (
	"fmt"

	"github.com/VarthanV/liv-server/fileservice"
	"github.com/gin-gonic/gin"
)

const (
	host = "127.0.0.1"
	port = 8060
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
	r.Run(fmt.Sprintf("%s:%d", host, port))
}
