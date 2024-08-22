package api

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
	"github.com/nmhoang2909/bank/token"
)

type Server struct {
	store         db.IStore
	route         *gin.Engine
	tokenMaker    token.Maker
	tokenDuration time.Duration
}

func NewServer(store db.IStore) (*Server, error) {
	maker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil, err
	}
	tokenDuration := time.Minute * 15

	sv := &Server{
		store:         store,
		tokenMaker:    maker,
		tokenDuration: tokenDuration,
	}

	sv.setRoutes()

	return sv, nil
}

func (s *Server) Start(address string) error {
	return s.route.Run(address)
}

func (s *Server) setRoutes() {
	router := gin.Default()
	router.POST("/accounts", s.createAccount)
	router.GET("/accounts/:id", s.getAccount)
	router.GET("/accounts", s.getAccounts)
	router.PUT("/accounts", s.updateBalanceAccount)
	router.DELETE("/accounts/:id", s.deleteAccount)

	router.POST("/users/login", s.userLogin)
	router.POST("/users", s.createUser)

	s.route = router
}
