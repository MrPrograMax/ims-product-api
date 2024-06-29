package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"ims-product-api/pkg/service"

	_ "ims-product-api/docs"
)

type Handler struct {
	product *service.ProductService
	auth    *service.AuthService
	order   *service.OrderService
	supply  *service.SupplyService
}

func NewHandler(product *service.ProductService, auth *service.AuthService, order *service.OrderService, supply *service.SupplyService) *Handler {
	return &Handler{
		product: product,
		auth: auth,
		order: order,
		supply: supply,
		}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	product := router.Group("/product")
	{
		product.POST("", h.addProduct)

		product.GET("/name/:name", h.findProductByName)
		product.GET("/id/:id", h.findProductById)

		product.GET("", h.getProducts)
		product.GET("/category_id/:id", h.findProductsByCategoryId)
		product.GET("/category_name/:name", h.findProductByCategoryName)
		product.GET("/loc", h.findProductByLocation)
		product.GET("/loc_id/:id", h.findProductsByLocationId)
		product.GET("/status_name/:name", h.findProductByStatusName)
		product.GET("/status_id/:id", h.findProductsByStatusId)

		product.PUT("/:id", h.updateProduct)

		product.DELETE("/:id", h.deleteProduct)

		category := product.Group("/category")
		{
			category.GET("", h.getCategories)
			category.GET("/name/:name", h.findCategoryByName)
			category.GET("/id/:id", h.findCategoryById)
		}

		location := product.Group("/location")
		{
			location.POST("", h.addLocation)

			location.GET("/row/:row/place/:place", h.findLocationByRowAndPlace)
			location.GET("/id/:id", h.findLocationById)

			location.GET("", h.getLocations)
			location.GET("/row/:row", h.findLocationByRow)
			location.DELETE("/:id", h.deleteLocation)
		}

		status := product.Group("/status")
		{
			status.GET("", h.getProductStatuses)
			status.GET("/name/:name", h.findProductStatusByName)
			status.GET("/id/:id", h.findProductStatusById)
		}
	}

	order := router.Group("/order")
	{
		order.POST("", h.addOrder)
		order.GET("", h.getOrders)
		order.DELETE("/:id", h.deleteOrder)

		order.POST("/item", h.addOrderItems)
		order.GET("/:id", h.getOrderDetails)
		order.GET("/product/:id", h.getProductOrders)
		order.DELETE("/:id/item/:item_id", h.deleteOrderItem)
	}

	supply := router.Group("/supply")
	{
		supply.POST("", h.addSupply)
		supply.GET("", h.getSupplies)
		supply.DELETE("/:id", h.deleteSupply)

		supply.POST("/item", h.addSupplyItems)
		supply.GET("/:id", h.getSupplyDetails)
		supply.GET("/product/:id", h.getProductSupplies)
		supply.DELETE("/item/:id", h.deleteSupplyItem)
	}

	return router
}
