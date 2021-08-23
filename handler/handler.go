package handler

import "github.com/gin-gonic/gin"

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
		lists := api.Group("/api/v1")
		{
			lists.POST("/employee_add", h.employee_add)
			lists.DELETE("/employee_remove", h.employee_remove)
			lists.PUT("/employee_upd", h.employee_upd)
			lists.GET("/get_all", h.get_all)
			lists.PUT("/employee_get", h.employee_get)

		}

	}

	return router
}
