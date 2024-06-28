package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func newStatusResponse(c *gin.Context, json any) {
	if json == nil {
		logrus.Info(http.StatusNoContent, json)
		c.Status(http.StatusNoContent)
		return
	}

	logrus.Info(http.StatusOK, json)
	c.JSON(http.StatusOK, json)
}
