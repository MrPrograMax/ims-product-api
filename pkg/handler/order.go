package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/model"
	"net/http"
)

// addOrder godoc
// @Summary Post order
// @Description Post a new order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} handler.idResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order [post]
func (h *Handler) addOrder(c *gin.Context) {
	id, err := h.order.Order.Create()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, idResponse{Id: id})
}

// getOrders godoc
// @Summary Get Orders
// @Description Get list of existing orders
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} model.Order
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order [get]
func (h *Handler) getOrders(c *gin.Context) {
	orders, err := h.order.Order.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, orders)
}

// deleteOrder godoc
// @Summary Delete order
// @Description Delete order by its id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order/{id} [delete]
func (h *Handler) deleteOrder(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.order.Order.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newStatusResponse(c, statusResponse{"ok"})
}

// addOrderItems godoc
// @Summary Post order item
// @Description Post a new product with info
// @Tags OrderItem
// @Accept json
// @Produce json
// @Param input body model.OrderItem true "OrderItem info"
// @Success 200 {array} handler.idResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order/item [post]
func (h *Handler) addOrderItems(c *gin.Context) {
	var items []model.OrderItem
	if err := c.BindJSON(&items); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ids, err := h.order.OrderItem.CreateList(items)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, toIdResponseSlice(ids))
}

// findLocationByRowAndPlace godoc
// @Summary Get order items by order id
// @Description Get list of order items for specified order by order id
// @Tags OrderItem
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {array} model.OrderItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order/{id} [get]
func (h *Handler) getOrderDetails(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	orderItems, err := h.order.OrderItem.GetByOrderId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, orderItems)
}

// getProductOrders godoc
// @Summary Get order items by product id
// @Description Get list of order items for specified product by product id
// @Tags OrderItem
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {array} model.OrderItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order/product/{id} [get]
func (h *Handler) getProductOrders(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	productOrders, err := h.order.OrderItem.GetByProductId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, productOrders)
}

// deleteOrderItem godoc
// @Summary Delete order item
// @Description Delete order item by its id
// @Tags OrderItem
// @Accept json
// @Produce json
// @Param id path int true "OrderItem ID"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /order/item/{id} [delete]
func (h *Handler) deleteOrderItem(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.order.OrderItem.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newStatusResponse(c, statusResponse{"ok"})
}
