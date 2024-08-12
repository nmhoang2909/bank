package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
)

type Server struct {
	store *db.Store
	route *gin.Engine
}

func NewServer(store *db.Store) *Server {
	sv := &Server{
		store: store,
	}
	router := gin.Default()
	router.POST("/accounts", sv.createAccount())
	sv.route = router
	return sv
}
