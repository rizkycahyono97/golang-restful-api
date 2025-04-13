package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

// constructor
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// CREATE
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category (name) VALUE (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err) // helper

	// auto increment
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	// memasukan ke id
	category.Id = int(id)
	return category
}

// UPDATE
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

// DELETE
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

// FindById
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{} // untuk penyimpanan sementara hasil query
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name) // untuk memasukan ke domain.Category{}
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

// FindAll
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category // untuk penyimpanan sementara hasil query
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
