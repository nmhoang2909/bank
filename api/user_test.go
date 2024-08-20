package api

import (
	"bytes"
	"encoding/json"
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

func checkBodyMatchUser(t *testing.T, body *bytes.Buffer, expect db.User) {
	bs, err := io.ReadAll(body)
	assert.NoError(t, err)

	type data struct {
		Data db.User `json:"data"`
	}
	var actual data
	err = json.Unmarshal(bs, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expect, actual.Data)
}

func TestUserAPI(t *testing.T) {
	pw := util.RandomString(5)
	hashedPw, _ := util.HashPassword(pw)
	user := db.User{
		Username:     util.RandomString(6),
		FullName:     util.RandomString(10),
		Email:        util.RandomEmail(),
		HashPassword: hashedPw,
	}
	tests := []struct {
		name          string
		req           createUserReq
		stubs         func(*mockdb.MockIStore)
		checkResponse func(t *testing.T, resp *httptest.ResponseRecorder)
	}{
		{
			name: "ok",
			req: createUserReq{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: pw,
			},
			stubs: func(mi *mockdb.MockIStore) {
				mi.EXPECT().CreateUser(gomock.Any(), gomock.Eq(user)).Return(int64(0), nil)
			},
			checkResponse: func(t *testing.T, resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, resp.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db := mockdb.NewMockIStore(ctrl)
			sv := NewServer(db)
			test.stubs(db)

			body, _ := json.Marshal(test.req)
			request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
			assert.NoError(t, err)
			resp := httptest.NewRecorder()

			sv.route.ServeHTTP(resp, request)
			test.checkResponse(t, resp)
		})
	}
}
