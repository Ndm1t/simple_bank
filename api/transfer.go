package api

import (
	db "bankingApp/db/sqlc"
	"bankingApp/token"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transferRequest struct {
	AccountFromId int64   `json:"accountFromId" binding:"required,min=1"`
	AccountToId   int64   `json:"accountToId" binding:"required,min=1"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Currency      string  `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, ok := server.validAccount(ctx, req.AccountFromId, req.Currency)
	if !ok {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Username != account.Owner {
		err := errors.New("It is forbidden to transfer money from account that does not belong to the authorized user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if _, ok = server.validAccount(ctx, req.AccountToId, req.Currency); !ok {
		return
	}

	arg := db.TransferTxParams{
		FromAccountId: req.AccountFromId,
		ToAccountId:   req.AccountToId,
		Amount:        req.Amount,
	}

	transferResult, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transferResult)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "There is no account with given id",
				"data":    accountId,
			})
			return db.Account{}, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return db.Account{}, false
	}
	if account.Currency != currency {
		err = fmt.Errorf("account [%d] currency mismatch: %v vs %v", accountId, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return db.Account{}, false
	}
	return account, true
}
