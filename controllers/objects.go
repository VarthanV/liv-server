package controllers

import "errors"

var (
	ErrPathIsADirectory = errors.New(
		"path is a directory, exact path of file is to be given to be served")
	ErrFileDoesNotExist = errors.New(
		"file does not exist")
)
