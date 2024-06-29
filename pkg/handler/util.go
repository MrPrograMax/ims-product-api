package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func validateInt(c *gin.Context, tag string) (int64, error) {
	id, ok := c.Get(tag)
	if !ok {
		return 0, fmt.Errorf("param %s not found", tag)
	}

	idStr, ok := id.(string)
	if !ok {
		return 0, fmt.Errorf("param %s of invalid type", tag)
	}

	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return idInt, fmt.Errorf("param %s of invalid type", tag)
	}

	return idInt, nil
}

func validateString(c *gin.Context, tag string) (string, error) {
	str, ok := c.Get(tag)
	if !ok {
		return "", fmt.Errorf("param %s not found", tag)
	}

	strAsserted, ok := str.(string)
	if !ok {
		return "", fmt.Errorf("param %s of invalid type", tag)
	}

	return strAsserted, nil
}
