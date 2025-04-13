package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate // package validate
}

// constructor
func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		validate:           validate,
		DB:                 DB,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	//validate
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	//untuk database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//mengubah web.CategoryCreateRequest ➝ domain.Category, karena repository hanya mengenal domain
	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	//validate
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	//untuk database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//update berdasarkan id
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) //not found error handling
	}

	category.Name = request.Name // yang di update

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	//untuk database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//cari berdasarkan Id baru dihapus
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) //not found error handling
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	//untuk database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) //not found error handling
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	//untuk database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
