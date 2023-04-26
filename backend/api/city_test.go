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

func TestGetCityByID(t *testing.T) {
	city := randomCity()

	testCases := []struct {
		name string
		cityID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			cityID: city.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetCity(gomock.Any(), gomock.Eq(city.ID)).Times(1).Return(city, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCity(t, recorder.Body, city)
			},
		},
		{
			name: "NotFound",
			cityID: city.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetCity(gomock.Any(), gomock.Eq(city.ID)).Times(1).Return(db.City{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			cityID: city.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetCity(gomock.Any(), gomock.Eq(city.ID)).Times(1).Return(db.City{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			cityID: 0,
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

		t.Run(tc.name, func(t *testing.T){
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/city/%d", tc.cityID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		}) 
	}
}

func randomCity() db.City {
	return db.City{
		ID: util.RandomInt(1, 1000),
		Name: util.RandomString(10),
		CreatedAt: time.Date(2023, time.April, 25, 11, 0, 0, 0, time.UTC),
	}
}

func requireBodyMatchCity(t *testing.T, body *bytes.Buffer, city db.City) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotCity db.City
	err = json.Unmarshal(data, &gotCity)
	require.NoError(t, err)
	require.Equal(t, city, gotCity)
}