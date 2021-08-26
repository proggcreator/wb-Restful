package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/proggcreator/wb-Restful/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		mylists := api.Group("/api/v1")
		{
			mylists.POST("/employee_add", h.employee_add)
			mylists.DELETE("/employee_remove", h.employee_remove)
			mylists.PUT("/employee_upd", h.employee_upd)
			mylists.GET("/get_all", h.get_all)
			mylists.PUT("/employee_get", h.employee_get)

		}

	}

	return router
}
