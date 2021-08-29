package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) employee_add(c *gin.Context) {
curEmpl:=
}

func (h *Handler) employee_remove(c *gin.Context) {

}
func (h *Handler) employee_upd(c *gin.Context) {

}
func (h *Handler) get_all(c *gin.Context) {

}
func (h *Handler) employee_get(c *gin.Context) {
	//parse id param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error get employee id"})
		return
	}
	//
	fmt.Fprint(c.Writer, "Полученный id:  ")
	fmt.Fprint(c.Writer, id)

}
