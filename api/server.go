package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
)

type Server struct {
	store db.IStore
	route *gin.Engine
}

func NewServer(store db.IStore) *Server {
	sv := &Server{
		store: store,
	}
	router := gin.Default()
	router.POST("/accounts", sv.createAccount)
	router.GET("/accounts/:id", sv.getAccount)
	router.GET("/accounts", sv.getAccounts)
	router.PUT("/accounts", sv.updateBalanceAccount)
	router.DELETE("/accounts/:id", sv.deleteAccount)

	router.POST("/users", sv.createUser)
	sv.route = router
	return sv
}

func (s *Server) Start(address string) error {
	return s.route.Run(address)
}
