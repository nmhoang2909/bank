package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
	"github.com/nmhoang2909/bank/util"
)

type createUserReq struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPw, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Username:     req.Username,
		FullName:     req.FullName,
		Email:        req.Email,
		HashPassword: hashedPw,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response(id))
}
