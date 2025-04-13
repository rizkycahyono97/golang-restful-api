package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"
)

func main() {

	//constructor
	db := app.NewDb()
	validate := validator.New()

	//constructor
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController) // router

	//httpserver
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router), // memakai middleware
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
