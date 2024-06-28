package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/pkg/service"
)

type Handler struct {
	product *service.ProductService
}

func NewHandler(services *service.ProductService) *Handler {
	return &Handler{product: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	product := router.Group("/product")
	{
		product.POST("/add", h.addProduct)

		product.GET("/name/:name", h.findProductByName)
		product.GET("/id/:id", h.findProductById)

		product.GET("/", h.getProducts)
		product.GET("/category/:id", h.findProductsByCategoryId)
		product.GET("/category/:name", h.findProductByCategoryName)
		product.GET("/loc", h.findProductByLocation)
		product.GET("/loc/:id", h.findProductsByLocationId)
		product.GET("/status/:name", h.findProductByStatusName)
		product.GET("/status/:id", h.findProductsByStatusId)

		product.PUT("/:id", h.updateProduct)

		product.DELETE("/:id", h.deleteProduct)

		category := product.Group("/category")
		{
			category.GET("/", h.getCategories)
			category.GET("/name/:name", h.findCategoryByName)
			category.GET("/id/:id", h.findCategoryById)
		}

		location := product.Group("/location")
		{
			location.POST("/add", h.addLocation)

			location.GET("/row/:row/place/:place", h.findLocationByRowAndPlace)
			location.GET("/id/:id", h.findLocationById)

			location.GET("/", h.getLocations)
			location.GET("/row/:row", h.findLocationByRow)

			product.DELETE("/:id", h.deleteLocation)
		}

		status := product.Group("/status")
		{
			status.GET("/", h.getProductStatuses)
			status.GET("/name/:name", h.findProductStatusByName)
			status.GET("/id/:id", h.findProductStatusById)
		}
	}

	return router
}
