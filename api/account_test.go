package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/nmhoang2909/bank/db/mock"
	db "github.com/nmhoang2909/bank/db/sqlc"
	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func randomAccount() db.Account {
	return db.Account{
		ID:       int64(util.RandomNumber(1, 1000)),
		Owner:    util.RandomString(6),
		Balance:  int32(util.RandomNumber(1000, 5000)),
		Currency: util.RandomCurrency(),
	}
}

func checkBodyMatch(t *testing.T, body *bytes.Buffer, expect db.Account) {
	bs, err := io.ReadAll(body)
	assert.NoError(t, err)

	type data struct {
		Data db.Account `json:"data"`
	}
	var actual data
	err = json.Unmarshal(bs, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expect, actual.Data)
}

func TestGetAccountAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockIStore(ctrl)
	server := NewServer(store)

	account := randomAccount()

	// stub
	store.EXPECT().
		GetAccountById(gomock.Any(), account.ID).
		Times(1).
		Return(account, nil)

	// server exec
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/accounts/%d", account.ID), nil)
	assert.NoError(t, err)
	recorder := httptest.NewRecorder()
	server.route.ServeHTTP(recorder, request)

	// assert
	assert.Equal(t, http.StatusOK, recorder.Code)
	checkBodyMatch(t, recorder.Body, account)
}
