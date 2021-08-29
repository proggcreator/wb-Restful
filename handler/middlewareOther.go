package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) checkJsonType(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-type")
	if contentType != "application/json" {
		logrus.Fatalf("Invalid content-type")

	}

}

func (h *Handler) acceptJsonOrXml(c *gin.Context) {
	listAccept := c.Request.Header["Accept"]
	fmt.Println("Ttttttt")
	fmt.Println(listAccept)
	//x := c.Request.Header["Accept"] if json or xml => coooool
	/*listAccept := c.Request.Header["Accept"]

	for _, cur := range listAccept {
		if (cur == "json") || (cur == "json") {

		}
	}

	logrus.Fatalf("Invalid accept header")
	*/
}
