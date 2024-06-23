package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/pkg/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//api := router.Group("/staff/") //h.userIdentity
	//{
	//	roles := api.Group("/roles")
	//	{
	//		roles.POST("/add", h.createRole)
	//		roles.GET("/", h.getRolesList)
	//		roles.GET("/:id", h.getRoleByName)
	//	}
	//}

	return router
}
