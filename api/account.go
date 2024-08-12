package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nmhoang2909/bank/db/sqlc"
)

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int32  `json:"balance"`
	Currency string `json:"currency"`
}

func (s *Server) createAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createAccount CreateAccountParams
		if err := ctx.ShouldBindJSON(&createAccount); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		id, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
			Owner:    createAccount.Owner,
			Balance:  createAccount.Balance,
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
}
