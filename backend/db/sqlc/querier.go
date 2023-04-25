// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	DeleteCategory(ctx context.Context, id int64) error
	GetCategory(ctx context.Context, id int64) (Category, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
