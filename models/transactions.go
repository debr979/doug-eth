package models

type GetTransactionByHashRequest struct {
	TxHash string `json:"txHash" from:"txHash" binding:"required"`
}

type TransactionLog struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}

type GetTransactionByHashResponse struct {
	TxHash string           `json:"tx_hahs"`
	From   string           `json:"from"`
	To     string           `json:"to"`
	Nonce  int64            `json:"nonce"`
	Data   string           `json:"data"`
	Value  string           `json:"value"`
	Logs   []TransactionLog `json:"logs"`
}
