package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/tomoki-yamamura/simple_bank/db/mock"
	db "github.com/tomoki-yamamura/simple_bank/db/sqlc"
	"github.com/tomoki-yamamura/simple_bank/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdb.NewMockStore(ctrl)
	mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(mockStore)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		ID: util.RandomInt(0, 1000),
		Owner: util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance: util.RandomMoney(),
	}
}