package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getProductStatuses(c *gin.Context) {
	statuses, err := h.product.ProductStatus.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, statuses)
}

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