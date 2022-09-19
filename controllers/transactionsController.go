package controllers

import (
	"doug/services"
	msg "doug/utils/errMsg"
	"github.com/gin-gonic/gin"
)

type transactions struct {
	baseController
}

var Transactions transactions

func (r *transactions) GetTransactionByHash(c *gin.Context) {
	txHash := c.Param("txHash")
	if txHash == "" {
		r.ResponseJSONWithCtx(c, msg.Messenger.Get(msg.ErrorParams))
		return
	}

	r.ResponseJSONWithCtx(c, services.Transactions.GetTransaction(txHash))
}
