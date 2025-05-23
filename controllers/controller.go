package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) InitRoutes(r *gin.Engine) {
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/:path", c.ServeFile)
	r.Run("127.0.0.1:8060")
}
