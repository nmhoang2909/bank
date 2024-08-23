package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	db "github.com/nmhoang2909/bank/db/sqlc"
	"github.com/nmhoang2909/bank/token"
)

type CreateAccountParams struct {
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var createAccount CreateAccountParams
	if err := ctx.ShouldBindJSON(&createAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	id, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    payload.Issuer,
		Balance:  0,
		Currency: createAccount.Currency,
	})
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1452, 1062:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response(account))
}

type getAccountsQuery struct {
	PageId   int32 `form:"page_id" binding:"min=1"`
	PageSize int32 `form:"page_size" binding:"min=5,max=10"`
}

func (s *Server) getAccounts(ctx *gin.Context) {
	var query getAccountsQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accounts, err := s.store.GetAccounts(ctx, db.GetAccountsParams{
		Limit:  query.PageSize,
		Offset: (query.PageId - 1) * query.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response(accounts))
}

type updateBalanceAccountReq struct {
	Amount int32 `json:"amount" binding:"required"`
	ID     int64 `json:"id" binding:"required,min=0"`
}

func (s *Server) updateBalanceAccount(ctx *gin.Context) {
	var req updateBalanceAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.UpdateAccountBalance(ctx, db.UpdateAccountBalanceParams{
		Amount: req.Amount,
		ID:     req.ID,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

type deleteAccountReq struct {
	Id int64 `uri:"id" binding:"min=1"`
}

func (s *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.DeleteAccountTx(ctx, req.Id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
