package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
)

type CreateAccountParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var createAccount CreateAccountParams
	if err := ctx.ShouldBindJSON(&createAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    createAccount.Owner,
		Balance:  0,
		Currency: createAccount.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := s.store.GetAccountById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
