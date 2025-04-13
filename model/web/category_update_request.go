package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,max=100,min=1" json:"name"`
}
