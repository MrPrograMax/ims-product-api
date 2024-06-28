package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/model"
	"net/http"
)

// addProduct godoc
// @Summary Post product
// @Description Post a new product with info
// @Tags Product
// @Accept json
// @Produce json
// @Param input body model.Product true "Product info"
// @Success 200 {integer} int
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product [post]
func (h *Handler) addProduct(c *gin.Context) {
	var product model.Product
	if err := c.BindJSON(&product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.product.Product.Create(product)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, map[string]interface{}{
		"id": id,
	})
}

// getLocations godoc
// @Summary Get Products
// @Description Get list of existing products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product [get]
func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.product.Product.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// findLocationByRowAndPlace godoc
// @Summary Get product by id
// @Description Get existing product by its id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/id/{id} [get]
func (h *Handler) findProductById(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.product.Product.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, product)
}

// findLocationByRowAndPlace godoc
// @Summary Get product by name
// @Description Get existing product by its name
// @Tags Product
// @Accept json
// @Produce json
// @Param name path string true "Product Name"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/name/{name} [get]
func (h *Handler) findProductByName(c *gin.Context) {
	name, err := validateString(c, "name")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.product.Product.GetByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, product)
}

// findLocationByRowAndPlace godoc
// @Summary Get products by category id
// @Description Get list of existing product by id of category belonging to product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/category_id/{id} [get]
func (h *Handler) findProductsByCategoryId(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.product.Product.GetByCategoryId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// findLocationByRowAndPlace godoc
// @Summary Get products by row and place
// @Description Get list of existing products by name of category belonging to product
// @Tags Product
// @Accept json
// @Produce json
// @Param name path string true "Category Name"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/category_name/{name} [get]
func (h *Handler) findProductByCategoryName(c *gin.Context) {
	name, err := validateString(c, "name")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.product.Product.GetByCategoryName(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// findLocationByRowAndPlace godoc
// @Summary Get products by location id
// @Description Get existing product by id of location belonging to product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/loc/{id} [get]
func (h *Handler) findProductsByLocationId(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.product.Product.GetByLocationId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// findLocationByRowAndPlace godoc
// @Summary Get products by location
// @Description Get list of existing products by location belonging to product
// @Tags Product
// @Accept json
// @Produce json
// @Param input body model.Location true "Location info"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/loc [get]
func (h *Handler) findProductByLocation(c *gin.Context) {
	var loc model.Location
	if err := c.BindJSON(&loc); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.product.Product.GetByLocation(loc)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// findLocationByRowAndPlace godoc
// @Summary Get products by status id
// @Description Get list of existing products by id of status belonging to product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Status ID"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/status/{id} [get]
func (h *Handler) findProductsByStatusId(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.product.Product.GetByStatusId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// findLocationByRowAndPlace godoc
// @Summary Get products by status name
// @Description Get list of existing products by id of status belonging to product
// @Tags Product
// @Accept json
// @Produce json
// @Param name path string true "Status Name"
// @Success 200 {array} model.Product
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/status_name/{name} [get]
func (h *Handler) findProductByStatusName(c *gin.Context) {
	name, err := validateString(c, "name")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.product.Product.GetByStatusName(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

// addLocation godoc
// @Summary Update product
// @Description Update existing product with new info
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param input body model.UpdateProduct true "Product Info"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var product model.UpdateProduct
	if err := c.BindJSON(&product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.product.Product.Update(id, product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newStatusResponse(c, statusResponse{"ok"})
}

// deleteLocation godoc
// @Summary Delete product
// @Description Delete product by its id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.product.Product.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newStatusResponse(c, statusResponse{"ok"})
}
