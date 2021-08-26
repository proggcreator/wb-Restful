package handler

import (
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type JsonError struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Code   string `json:"code`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func NewJsonError(c *gin.Context, curerr JsonError, intcode int) {
	logrus.Error(curerr)
	//stop next handlers, write responce
	c.AbortWithStatusJSON(intcode, curerr)
}
