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

func TestGetCategoryByID(t *testing.T) {
	category := randomCategory()

	testCases := []struct {
		name string
		categoryID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			categoryID: category.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stub. This can be translated as: I expect the GetCategory function of the store to be called with any context and this specific account ID arguments
				// we can also specify how many times this function should be called using the Times() function
				// we can use the Return() function to tell gomock to return some specific values whenever the GetAccount() function is called
				store.EXPECT().GetCategory(gomock.Any(), gomock.Eq(category.ID)).Times(1).Return(category, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response http status code
				// in the happy case, it should be http.StatusOK, this status code is recorded in the Code field of the recorder
				require.Equal(t, http.StatusOK, recorder.Code)
				// check response body should match category object
				requireBodyMatchCategory(t, recorder.Body, category)
			},
		},
		{
			name: "NotFound",
			categoryID: category.ID,
			buildStubs: func (store *mockdb.MockStore){
				store.EXPECT().GetCategory(gomock.Any(), gomock.Eq(category.ID)).Times(1).Return(db.Category{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			categoryID: category.ID,
			buildStubs: func (store *mockdb.MockStore) {
				store.EXPECT().GetCategory(gomock.Any(), gomock.Eq(category.ID)).Times(1).Return(db.Category{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			categoryID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetCategory(gomock.Any(), gomock.Any()).Times(0)
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

			url := fmt.Sprintf("/category/%d", tc.categoryID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// send our API request through the server router and record its response in the recorder, then all we need to do is to check that response
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateCategory(t *testing.T) {
	category := randomCategory()
	
	testCases := []struct {
		name string
		body gin.H
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"title": category.Title,
				"description": category.Description,
				"image_url": category.ImageUrl,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateCategoryParams{
					Title: category.Title,
					Description: category.Description,
					ImageUrl: category.ImageUrl,
				}
				store.EXPECT().CreateCategory(gomock.Any(), gomock.Eq(arg)).Times(1).Return(category, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCategory(t, recorder.Body, category)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"title": category.Title,
				"description": category.Description,
				"image_url": category.ImageUrl,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateCategoryParams{
					Title: category.Title,
					Description: category.Description,
					ImageUrl: category.ImageUrl,
				}
				store.EXPECT().CreateCategory(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Category{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			// marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/category"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListCategory(t *testing.T) {
	n := 5
	categories := make([]db.Category, n)
	for i := 0; i < n; i++ {
		categories[i] = randomCategory()
	}

	testCases := []struct {
		name string
		query listCategoryRequest
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: listCategoryRequest{
				PageID: 1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListCategoriesParams{
					Limit: int32(n),
					Offset: 0,
				}
				store.EXPECT().ListCategories(gomock.Any(), gomock.Eq(arg)).Times(1).Return(categories, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCategories(t, recorder.Body, categories)
			},
		},
		{
			name: "InternalError",
			query: listCategoryRequest{
				PageID: 1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListCategoriesParams{
					Limit: int32(n),
					Offset: 0,
				}
				store.EXPECT().ListCategories(gomock.Any(), gomock.Eq(arg)).Times(1).Return([]db.Category{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: listCategoryRequest{
				PageID: -1, 
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListCategories(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: listCategoryRequest{
				PageID: 1,
				PageSize: 50,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListCategories(gomock.Any(), gomock.Any()).Times(0)
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

			url := "/category"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.PageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.PageSize))
			request.URL.RawQuery = q.Encode()

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomCategory() db.Category {
	return db.Category{
		ID: util.RandomInt(1, 1000),
		Title: util.RandomString(20),
		Description: util.RandomString(50),
		ImageUrl: util.RandomString(10),
		CreatedAt: time.Date(2023, time.April, 25, 11, 0, 0, 0, time.UTC),
	}
}

func requireBodyMatchCategory(t *testing.T, body *bytes.Buffer, category db.Category) {
	// call io.ReadAll() to read all data from the response body and store in data variable
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	// declare new gotCategory variable to store the category object we got from the response body data
	var gotCategory db.Category
	err = json.Unmarshal(data, &gotCategory)
	require.NoError(t, err)
	require.Equal(t, category, gotCategory)
}

func requireBodyMatchCategories(t *testing.T, body *bytes.Buffer, categories []db.Category) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotCategories []db.Category
	err = json.Unmarshal(data, &gotCategories)
	require.NoError(t, err)
	require.Equal(t, categories, gotCategories)
}