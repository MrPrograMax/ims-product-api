package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/model"
	"net/http"
)

func (h *Handler) addLocation(c *gin.Context) {
	var location model.Location
	if err := c.BindJSON(&location); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.product.Location.Create(location)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getLocations(c *gin.Context) {
	locations, err := h.product.Location.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, locations)
}

func (h *Handler) findLocationByRow(c *gin.Context) {
	row, err := validateString(c, "row")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	locations, err := h.product.Location.GetByRow(row)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, locations)
}

func (h *Handler) findLocationById(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.product.Location.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, product)
}

func (h *Handler) findLocationByRowAndPlace(c *gin.Context) {
	row, err := validateString(c, "row")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	place, err := validateString(c, "place")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	location, err := h.product.Location.GetByRowAndPlace(row, place)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, location)
}

func (h *Handler) deleteLocation(c *gin.Context) {
	id, err := validateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.product.Location.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	newStatusResponse(c, map[string]interface{}{"status": "ok"})
}
