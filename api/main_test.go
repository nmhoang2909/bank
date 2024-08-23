package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

func newTestServer(t *testing.T, store db.IStore) *Server {
	config, err := util.LoadConfig("./..")
	assert.NoError(t, err)

	sv, err := NewServer(store, config)
	assert.NoError(t, err)
	return sv
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
