package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getCategories godoc
// @Summary Get all categories
// @Description Get list of categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} model.Category
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/category [get]
func (h *Handler) getCategories(c *gin.Context) {
	cats, err := h.product.Category.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, cats)
}

// findCategoryById godoc
// @Summary Get category by id
// @Description Get existing category by its id
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/category/id/{id} [get]
func (h *Handler) findCategoryById(c *gin.Context) {
	id, err := validateInt(c, "id")
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

// findCategoryByName godoc
// @Summary Get categories by name
// @Description Get existing category by its name
// @Tags Category
// @Accept json
// @Produce json
// @Param name path string true "Category Name"
// @Success 200 {object} model.Category
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/category/name/{name} [get]
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
