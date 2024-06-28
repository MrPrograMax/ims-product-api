package model

type ProductDTO struct {
	Id          int64         `json:"id" db:"id"`
	Name        string        `json:"name" binding:"required"`
	Quantity    int64         `json:"quantity" binding:"required"`
	Description string        `json:"description" binding:"required"`
	Category    Category      `json:"category" binding:"required"`
	Location    Location      `json:"location" binding:"required"`
	Status      ProductStatus `json:"status" binding:"required"`
}

func (dto *ProductDTO) ToProduct() Product {
	return Product{
		Id:          dto.Id,
		Name:        dto.Name,
		Quantity:    dto.Quantity,
		Description: dto.Description,
		CategoryId:  dto.Category.Id,
		LocationId:  dto.Location.Id,
		StatusId:    dto.Status.Id,
	}
}
