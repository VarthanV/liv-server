package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *Controller) ServeFile(ctx *gin.Context) {
	path := ctx.Param("path")
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error("error in getting working directory ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	absolutePath := fmt.Sprintf("%s%s", pwd, path)
	logrus.Info("file path ", absolutePath)
	fileInfo, err := os.Stat(absolutePath)
	if err != nil {
		logrus.Error("error in checking if file can be served ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if fileInfo.IsDir() {
		files, err := c.fileService.List(absolutePath, path)
		if err != nil {
			logrus.Error("error in listing files ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.HTML(http.StatusOK, "list_files.html", gin.H{
			"DirectoryName": path,
			"Files":         files,
		})
		return
	}
	ctx.File(absolutePath)
}
