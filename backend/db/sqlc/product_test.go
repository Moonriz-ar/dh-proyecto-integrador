package db

import (
	"context"
	"database/sql"
	"proyecto-integrador/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Title, product2.Title)
	require.Equal(t, product1.Description, product2.Description)
	require.Equal(t, product1.CategoryID, product2.CategoryID)
	require.Equal(t, product1.CityID, product2.CityID)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestListProduct(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductParams{
		Limit: 5,
		Offset: 5,
	}

	products, err := testQueries.ListProduct(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	arg := UpdateProductParams{
		ID: product1.ID,
		Title: product1.Title,
		Description: product1.Description,
		CategoryID: product1.CategoryID,
		CityID: product1.CityID,
	}

	product2, err := testQueries.UpdateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, arg.ID, product2.ID)
	require.Equal(t, arg.Title, product2.Title)
	require.Equal(t, arg.Description, product2.Description)
	require.Equal(t, arg.CategoryID, product2.CategoryID)
	require.Equal(t, arg.CityID, product2.CityID)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestDeleteProdcut(t *testing.T) {
	product1 := createRandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Title: util.RandomString(10),
		Description: util.RandomString(30),
		CategoryID: util.RandomInt(1, 6), // in sql init script, 6 categories are inserted
		CityID: util.RandomInt(1, 30), // in sql init script, 30 argentina cities are inserted
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)
	
	require.Equal(t, arg.Title, product.Title)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.Equal(t, arg.CityID, product.CityID)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}