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
		},{
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
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/category/%d", tc.categoryID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// send our API request through the server router and record its response in the recorder, then all we need to do is to check that response
			server.router.ServeHTTP(recorder, request)
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