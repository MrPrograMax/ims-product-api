package handler

import (
	"github.com/gin-gonic/gin"
	"ims-product-api/model"
	"net/http"
)

// addLocation godoc
// @Summary Post location
// @Description Post a new location with info
// @Tags Location
// @Accept json
// @Produce json
// @Param input body model.Location true "Location info"
// @Success 200 {integer} int
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/location [post]
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

// getLocations godoc
// @Summary Get locations
// @Description Get list of existing locations
// @Tags Location
// @Accept json
// @Produce json
// @Success 200 {array} model.Location
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/location [get]
func (h *Handler) getLocations(c *gin.Context) {
	locations, err := h.product.Location.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newStatusResponse(c, locations)
}

// findLocationByRow godoc
// @Summary Get locations by row
// @Description Get list of existing locations by its row
// @Tags Location
// @Accept json
// @Produce json
// @Param row path string true "Location Row"
// @Success 200 {array} model.Location
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/location/row/{row} [get]
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

// findLocationById godoc
// @Summary Get status by id
// @Description Get existing location by its id
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Success 200 {object} model.Location
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/location/id/{id} [get]
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

// findLocationByRowAndPlace godoc
// @Summary Get location by row and place
// @Description Get existing location by its row and place
// @Tags Location
// @Accept json
// @Produce json
// @Param row path string true "Location Row"
// @Param place path string true "Location Place"
// @Success 200 {object} model.Location
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/location/row/{row}/place/{place} [get]
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

// deleteLocation godoc
// @Summary Delete location
// @Description Delete location by its id
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Success 200 {object} handler.statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /product/location/{id} [delete]
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

	newStatusResponse(c, statusResponse{"ok"})
}
