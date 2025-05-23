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
	absolutePath := fmt.Sprintf("%s/%s", pwd, path)
	logrus.Info("file path ", absolutePath)
	fileInfo, err := c.canServe(absolutePath)
	if err != nil {
		logrus.Error("error in checking if file can be served ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("File-Name", fileInfo.Name())
	ctx.File(absolutePath)
}

func (c *Controller) canServe(path string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileDoesNotExist
		}
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, ErrPathIsADirectory
	}
	return fileInfo, nil
}
