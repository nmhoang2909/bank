package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nmhoang2909/bank/token"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "authorized",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				token, err := tokenMaker.CreateToken("user", time.Minute)
				assert.NoError(t, err)
				req.Header.Set(authorizationHeaderKey, authorizationTypeBearer+" "+token)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "unauthorization due to invalid bearer type",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				token, err := tokenMaker.CreateToken("user", time.Minute)
				assert.NoError(t, err)
				req.Header.Set(authorizationHeaderKey, "beeee "+token)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "unauthorization due to invalid token",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				req.Header.Set(authorizationHeaderKey, authorizationTypeBearer+" faketoken")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "unauthorization due to empty header",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				req.Header.Set(authorizationHeaderKey, "")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "unauthorization due to invalid type header format",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				req.Header.Set(authorizationHeaderKey, "Bearer ")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "unauthorization due to token is expired",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				token, err := tokenMaker.CreateToken("user", 2*-time.Minute)
				assert.NoError(t, err)
				req.Header.Set(authorizationHeaderKey, authorizationTypeBearer+" "+token)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			server.route.GET("/auth", authMiddleware(server.tokenMaker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			req, err := http.NewRequest(http.MethodGet, "/auth", nil)
			assert.NoError(t, err)
			recorder := httptest.NewRecorder()
			test.setupAuth(t, req, server.tokenMaker)
			server.route.ServeHTTP(recorder, req)
			test.checkResponse(t, recorder)
		})
	}
}
