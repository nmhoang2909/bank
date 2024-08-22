package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

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

	params := db.CreateUserParams{
		Username:     req.Username,
		FullName:     req.FullName,
		Email:        req.Email,
		HashPassword: hashedPw,
	}

	id, err := s.store.CreateUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response(id))
}

type userLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type userLoginResp struct {
	AccessToken string   `json:"access_token"`
	User        userResp `json:"user"`
}

type userResp struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Server) userLogin(ctx *gin.Context) {
	var req userLoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = util.IsCorrectPassword([]byte(user.HashPassword), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Username, s.tokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := userLoginResp{
		AccessToken: token,
		User: userResp{
			Username:  user.Username,
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}

	ctx.JSON(http.StatusOK, response(resp))
}
