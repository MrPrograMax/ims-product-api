package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/model"
	"net/http"
)

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

func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.product.Product.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, products)
}

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

	newStatusResponse(c, map[string]interface{}{"status": "ok"})
}

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

	newStatusResponse(c, map[string]interface{}{"status": "ok"})
}
