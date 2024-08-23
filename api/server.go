package api

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
	"github.com/nmhoang2909/bank/token"
	"github.com/nmhoang2909/bank/util"
)

type Server struct {
	store         db.IStore
	route         *gin.Engine
	tokenMaker    token.Maker
	tokenDuration time.Duration
	config        util.Config
}

func NewServer(store db.IStore, config util.Config) (*Server, error) {
	maker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, err
	}
	tokenDuration := time.Minute * 15

	sv := &Server{
		store:         store,
		tokenMaker:    maker,
		tokenDuration: tokenDuration,
		config:        config,
	}

	sv.setRoutes()

	return sv, nil
}

func (s *Server) Start(address string) error {
	return s.route.Run(address)
}

func (s *Server) setRoutes() {
	router := gin.Default()
	router.POST("/users/login", s.userLogin)
	router.POST("/users", s.createUser)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	authRoutes.POST("/accounts", s.createAccount)
	authRoutes.GET("/accounts/:id", s.getAccount)
	authRoutes.GET("/accounts", s.getAccounts)
	authRoutes.PUT("/accounts", s.updateBalanceAccount)
	authRoutes.DELETE("/accounts/:id", s.deleteAccount)

	s.route = router
}
