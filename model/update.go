package model

import (
	"fmt"
	"reflect"
)

type UpdateProduct struct {
	Name        *string `json:"name"`
	Quantity    *int64  `json:"quantity"`
	Description *string `json:"description"`
	CategoryId  *int64  `json:"category_id"`
	LocationId  *int64  `json:"location_id"`
	StatusId    *int64  `json:"status_id"`
}

type UpdateLocation struct {
	Row   *string `json:"row"`
	Place *string `json:"place"`
}

type UpdateCategory struct {
	Name *string `json:"name" binding:"required"`
}

func Verify(item interface{}) error {
	prodType := reflect.ValueOf(item).Elem()
	for i := 0; i < prodType.NumField(); i++ {
		valueField := prodType.Field(i)
		if valueField.Kind() == reflect.Interface {
			continue
		}

		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			return nil
		}
	}
	return fmt.Errorf("updating entity is empty")
}
