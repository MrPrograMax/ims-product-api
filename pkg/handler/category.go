package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCategories(c *gin.Context) {
	cats, err := h.product.Category.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, cats)
}

func (h *Handler) findCategoryById(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.product.Category.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, category)
}

func (h *Handler) findCategoryByName(c *gin.Context) {
	name, err := validateString(c, "name")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.product.Category.GetByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, category)
}
