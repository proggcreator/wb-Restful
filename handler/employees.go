package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	restful "github.com/proggcreator/wb-Restful"
)

func (h *Handler) employee_add(c *gin.Context) {
	//create new employee
	exampleEmp := restful.RetEmployee()
	//
	id, err := h.services.EmplWork.CreateEmpl(exampleEmp)
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error create employee"})
		return
	}
	fmt.Fprint(c.Writer, "Создан пользователь с id:  ")
	fmt.Fprint(c.Writer, id)
}

func (h *Handler) employee_remove(c *gin.Context) {

}
func (h *Handler) employee_upd(c *gin.Context) {

}
func (h *Handler) get_all(c *gin.Context) {
	list, err := h.services.EmplWork.GetAllEmpl()
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error create employee"})
		return
	}
	fmt.Fprint(c.Writer, "Все пользователи:  ")
	fmt.Fprint(c.Writer, list)
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
