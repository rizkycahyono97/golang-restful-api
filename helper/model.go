package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
)

// Konversi dari domain.Category ke web.CategoryResponse
// Tujuannya untuk mengubah struct internal dari domain Category menjadi struct khusus untuk output API (web.CategoryResponse).
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// Konversi dari domain.Category ke web.CategoryResponse
// bertugas untuk mengubah setiap elemen dalam slice domain.Category menjadi web.CategoryResponse.
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
