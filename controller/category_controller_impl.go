package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

// polymorpishm/constructor
func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Decode JSON Request Body
	categoryCreateRequest := web.CategoryCreateRequest{} // menyimpan struct
	helper.ReadFromRequest(r, &categoryCreateRequest)

	//Panggil Service Layer
	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	//Siapkan Response JSON
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	//Kirimkan JSON Response ke Client
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Decode JSON Request Body
	categoryUpdateRequest := web.CategoryUpdateRequest{} // menyimpan struct
	helper.ReadFromRequest(r, &categoryUpdateRequest)

	//Ambil ID dari URL dan Konversi ke Integer
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	//Panggil Service Layer
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	//Siapkan Response JSON
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	//Kirimkan JSON Response ke Client
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Ambil ID dari URL dan Konversi ke Integer
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	//Panggil Service Layer
	controller.CategoryService.Delete(r.Context(), id)

	//Siapkan Response JSON
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	//Kirimkan JSON Response ke Client
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Ambil ID dari URL dan Konversi ke Integer
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	//Panggil Service Layer
	categoryResponse := controller.CategoryService.FindById(r.Context(), id)

	//Siapkan Response JSON
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	//Kirimkan JSON Response ke Client
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Panggil Service Layer
	categoryResponses := controller.CategoryService.FindAll(r.Context())

	//Siapkan Response JSON
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	//Kirimkan JSON Response ke Client
	helper.WriteToResponseBody(w, webResponse)
}
