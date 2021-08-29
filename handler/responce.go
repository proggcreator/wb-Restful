package handler

import (
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type JsonError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func NewJsonError(c *gin.Context, curerr JsonError) {
	logrus.Error(curerr)
	//stop next handlers, write responce
	c.AbortWithStatusJSON(curerr.Status, curerr)
}
