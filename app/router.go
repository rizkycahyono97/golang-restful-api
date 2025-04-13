package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/controller"
	"golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	//inisialisasi package httprouter
	router := httprouter.New()

	//router
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	//error handling
	router.PanicHandler = exception.ErrorHandler

	return router
}
