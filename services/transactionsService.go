package services

import (
	msg "doug/utils/errMsg"
	"doug/utils/ether"
	"strings"
)

type transactions struct {
}

var Transactions transactions

func (r *transactions) GetTransaction(txHash string) interface{} {
	if !strings.Contains(txHash, "0x") {
		return msg.Messenger.Get(msg.ErrorParams)
	}

	return ether.Eth.GetTransactionByHash(txHash)
}
