package dto

type CreateCategoryRequest struct {
	Name         string `json:"name"`
	ParentId     uint   `json:"parentid"`
	ImageUrl     string `json:"imageurl"`
	DisplayOrder int    `json:"displayorder"`
}
