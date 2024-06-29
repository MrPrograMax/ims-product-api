package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/model"
	"net/http"
)

// addSupply godoc
// @Summary Post supply
// @Description Post a new supply
// @Tags Supply
// @Accept json
// @Produce json
// @Success 200 {object} handler.idResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply [post]
func (h *Handler) addSupply(c *gin.Context) {
	id, err := h.supply.Supply.Create()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, idResponse{Id: id})
}

// getSupplies godoc
// @Summary Get Supplies
// @Description Get list of existing supplies
// @Tags Supply
// @Accept json
// @Produce json
// @Success 200 {array} model.Supply
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply [get]
func (h *Handler) getSupplies(c *gin.Context) {
	supplies, err := h.supply.SupplyItem.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, supplies)
}

// deleteSupply godoc
// @Summary Delete supply
// @Description Delete supply by its id
// @Tags Supply
// @Accept json
// @Produce json
// @Param id path int true "Supply ID"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply/{id} [delete]
func (h *Handler) deleteSupply(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.supply.Supply.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newStatusResponse(c, statusResponse{"ok"})
}

// addSupplyItems godoc
// @Summary Post supply item
// @Description Post a new product with info
// @Tags SupplyItem
// @Accept json
// @Produce json
// @Param input body model.SupplyItem true "SupplyItem info"
// @Success 200 {array} handler.idResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply/item [post]
func (h *Handler) addSupplyItems(c *gin.Context) {
	var items []model.SupplyItem
	if err := c.BindJSON(&items); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ids, err := h.supply.SupplyItem.CreateList(items)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, toIdResponseSlice(ids))
}

// findLocationByRowAndPlace godoc
// @Summary Get supply items by supply id
// @Description Get list of supply items for specified supply by supply id
// @Tags SupplyItem
// @Accept json
// @Produce json
// @Param id path int true "Supply ID"
// @Success 200 {array} model.SupplyItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply/{id} [get]
func (h *Handler) getSupplyDetails(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	supplyItems, err := h.supply.SupplyItem.GetBySupplyId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, supplyItems)
}

// getProductSupplies godoc
// @Summary Get supply items by product id
// @Description Get list of supply items for specified product by product id
// @Tags SupplyItem
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {array} model.SupplyItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply/product/{id} [get]
func (h *Handler) getProductSupplies(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	productSupplies, err := h.supply.SupplyItem.GetByProductId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, productSupplies)
}

// deleteSupplyItem godoc
// @Summary Delete supply item
// @Description Delete supply item by its id
// @Tags SupplyItem
// @Accept json
// @Produce json
// @Param id path int true "SupplyItem ID"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /supply/item/{id} [delete]
func (h *Handler) deleteSupplyItem(c *gin.Context) {
	id, err := validateInt(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.supply.SupplyItem.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newStatusResponse(c, statusResponse{"ok"})
}
