package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) checkJsonType(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-type")
	if contentType != "application/json" {
		logrus.Fatalf("Invalid content-type")

	}
	//x := c.Request.Header["Accept"] if json or xml => coooool
}

func (h *Handler) acceptJsonOrXml(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-type")
	if contentType != "application/json" {
		logrus.Fatalf("Invalid content-type")

	}
}
