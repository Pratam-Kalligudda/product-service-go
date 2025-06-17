package dto

type UpdateProductDTO struct {
	Name        *string
	Description *string
	Stock       *int
	Price       *float64
	CategoryID  *uint
}

type CategoryDTO struct {
	Name string
}
