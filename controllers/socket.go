package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins; restrict this in production!
	},
}

func (c *Controller) HandleSocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Error("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Info("Read error:", err)
			break
		}
		logrus.Infof("Received: %s\n", message)

		// Echo back the message
		if err := conn.WriteMessage(messageType, message); err != nil {
			logrus.Info("Write error:", err)
			break
		}
		c.fileService.InitWatcher(ctx, conn)
	}
}
