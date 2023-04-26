package db

import (
	"context"
	"proyecto-integrador/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateCity(t *testing.T) {
	createRandomCity(t)
}

func TestGetCity(t *testing.T) {
	city1 := createRandomCity(t)
	city2, err := testQueries.GetCity(context.Background(), city1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, city2)

	require.Equal(t, city1.ID, city2.ID)
	require.Equal(t, city1.Name, city2.Name)
	require.WithinDuration(t, city1.CreatedAt, city2.CreatedAt, time.Second)
}

func ListCities(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCity(t)
	}

	arg := ListCitiesParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListCities(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createRandomCity(t *testing.T) City {
	arg := struct {
		Name string
	}{
		Name: util.RandomString(10),
	}
	city, err := testQueries.CreateCity(context.Background(), arg.Name)

	require.NoError(t, err)
	require.NotEmpty(t, city)

	require.Equal(t, arg.Name, city.Name)

	require.NotZero(t, city.ID)
	require.NotZero(t, city.CreatedAt)

	return city
}