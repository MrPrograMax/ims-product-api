package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getProductStatuses godoc
// @Summary Get all statuses
// @Description Get list of statuses
// @Tags ProductStatus
// @Accept json
// @Produce json
// @Success 200 {array} model.ProductStatus
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/status [get]
func (h *Handler) getProductStatuses(c *gin.Context) {
	statuses, err := h.product.ProductStatus.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, statuses)
}

// findProductStatusById godoc
// @Summary Get status by id
// @Description Get existing status by its id
// @Tags ProductStatus
// @Accept json
// @Produce json
// @Param id path int true "ProductStatus ID"
// @Success 200 {object} model.ProductStatus
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/status/id/{id} [get]
func (h *Handler) findProductStatusById(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.product.ProductStatus.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, status)
}

// findProductStatusByName godoc
// @Summary Get status by name
// @Description Get existing status by its name
// @Tags ProductStatus
// @Accept json
// @Produce json
// @Param name path string true "ProductStatus Name"
// @Success 200 {object} model.ProductStatus
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/status/name/{name} [get]
func (h *Handler) findProductStatusByName(c *gin.Context) {
	name, err := validateString(c, "name")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.product.ProductStatus.GetByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, status)
}
