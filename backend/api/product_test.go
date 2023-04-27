package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "proyecto-integrador/db/mock"
	db "proyecto-integrador/db/sqlc"
	"proyecto-integrador/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	product := randomProduct()

	testCases := []struct{
		name string
		body gin.H
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H {
				"title": product.Title,
				"description": product.Description,
				"category_id": product.CategoryID,
				"city_id": product.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateProductParams{
					Title: product.Title,
					Description: product.Description,
					CategoryID: product.CategoryID,
					CityID: product.CityID,
				}
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"title": product.Title,
				"description": product.Description,
				"category_id": product.CategoryID,
				"city_id": product.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateProductParams{
					Title: product.Title,
					Description: product.Description,
					CategoryID: product.CategoryID,
					CityID: product.CityID,
				}
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidRequestBody",
			body: gin.H{
				"title": product.Title,
				"category_id": product.CategoryID,
				"city_id": product.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/product"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	product := randomProduct()

	testCases := []struct{
		name string
		productID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name: "NotFound",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(db.Product{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().GetProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			productID: -1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases{
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/product/%d", tc.productID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListProduct(t *testing.T) {
	n := 5
	products := make([]db.Product, n)
	for i := 0; i < n; i++ {
		products[i] = randomProduct()
	}

	testCases := []struct{
		name string
		query listProductRequest
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: listProductRequest{
				PageID: 1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListProductParams{
					Limit: int32(n),
					Offset: 0,
				}
				store.EXPECT().ListProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(products, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProducts(t, recorder.Body, products)
			},
		},
		{
			name: "InternalError",
			query: listProductRequest{
				PageID: 1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListProductParams{
					Limit: int32(n),
					Offset: 0,
				}
				store.EXPECT().ListProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return([]db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: listProductRequest{
				PageID: -1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: listProductRequest{
				PageID: 1,
				PageSize: 50,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := "/product"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// add query parameters to request url
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.PageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.PageSize))
			request.URL.RawQuery = q.Encode()

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateProductByID(t *testing.T) {
	product1 := randomProduct()
	product2 := randomProduct()

	testCases := []struct {
		name string
		productID int64
		body gin.H
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			productID: product1.ID,
			body: gin.H{
				"title": product2.Title,
				"description": product2.Description,
				"category_id": product2.CategoryID,
				"city_id": product2.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product1.ID,
					Title: product2.Title,
					Description: product2.Description,
					CategoryID: product2.CategoryID,
					CityID: product2.CityID,
				}
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(product2, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product2)
			},
		},
		{
			name: "InternalError",
			productID: product1.ID,
			body: gin.H {
				"title": product2.Title,
				"description": product2.Description,
				"category_id": product2.CategoryID,
				"city_id": product2.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product1.ID,
					Title: product2.Title,
					Description: product2.Description,
					CategoryID: product2.CategoryID,
					CityID: product2.CityID,
				}
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NotFound",
			productID: product1.ID,
			body: gin.H {
				"title": product2.Title,
				"description": product2.Description,
				"category_id": product2.CategoryID,
				"city_id": product2.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product1.ID,
					Title: product2.Title,
					Description: product2.Description,
					CategoryID: product2.CategoryID,
					CityID: product2.CityID,
				}
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Product{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			productID: -1,
			body: gin.H {
				"title": product2.Title,
				"description": product2.Description,
				"category_id": product2.CategoryID,
				"city_id": product2.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidBody",
			productID: product1.ID,
			body: gin.H {
				"title": product2.Title,
				"city_id": product2.CityID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/product/%d", tc.productID)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDeleteProductByID(t *testing.T) {
	product := randomProduct()

	testCases := []struct{
		name string
		productID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name: "InternalError",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NotFound",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			productID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/product/%d", tc.productID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomProduct() db.Product {
	return db.Product{
		ID: util.RandomInt(1, 1000),
		Title: util.RandomString(20),
		Description: util.RandomString(30),
		CategoryID: util.RandomInt(1, 6),
		CityID: util.RandomInt(1,30),
		CreatedAt: time.Date(2023, time.April, 25, 11, 0, 0, 0, time.UTC),
	}
}

func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Product) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProduct db.Product
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, product, gotProduct)
}

func requireBodyMatchProducts(t *testing.T, body *bytes.Buffer, products []db.Product) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProducts []db.Product
	err = json.Unmarshal(data, &gotProducts)
	require.NoError(t, err)
	require.Equal(t, products, gotProducts)
}