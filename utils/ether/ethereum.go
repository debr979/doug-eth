package ether

import (
	"context"
	"doug/models"
	lg "doug/utils/logging"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type eth struct {
	client  *ethclient.Client
	chainId *big.Int
}

var Eth eth

func (r *eth) ClientInit() {
	ctx := context.Background()
	client, err := ethclient.Dial(os.Getenv("RCP_URL_PROD"))
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}
	r.client = client
	r.chainId, err = client.ChainID(ctx)
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}
}

func (r *eth) GetBlockNumber() {
	var err error
	frequency, err := strconv.Atoi(os.Getenv("SYNC_FREQ"))
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}

	ctx := context.Background()
	var blockNum uint64
	for {
		blockNum, err = r.client.BlockNumber(ctx)
		if err != nil {
			lg.Logging.ErrorLogging(err)
		}

		log.Println(blockNum)
		time.Sleep(time.Duration(frequency) * time.Millisecond)
	}
}

func (r *eth) GetBlockInfoByCount(blocksCount int) []models.BlockInfo {
	ctx := context.Background()

	blockNum, err := r.client.BlockNumber(ctx)
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}

	blockInfos := make([]models.BlockInfo, 0)
	for i := blockNum - uint64(blocksCount); i < blockNum; i++ {
		info, err := r.client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			lg.Logging.ErrorLogging(err)
		}

		var blockInfo models.BlockInfo
		blockInfo.BlockNum = info.Number().Int64()
		blockInfo.BlockHash = info.Hash().String()
		blockInfo.BlockTime = int64(info.Time())
		blockInfo.ParentHash = info.ParentHash().String()
		blockInfos = append(blockInfos, blockInfo)
	}

	return blockInfos
}

func (r *eth) GetBlockInfoById(blockNum int64) models.GetBlockByIdResponse {
	ctx := context.Background()

	info, err := r.client.BlockByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}

	var blockInfo models.GetBlockByIdResponse
	blockInfo.BlockNum = info.Number().Int64()
	blockInfo.BlockHash = info.Hash().String()
	blockInfo.BlockTime = int64(info.Time())
	blockInfo.ParentHash = info.ParentHash().String()
	blockInfo.Transactions = make([]string, 0)
	for _, t := range info.Transactions() {
		blockInfo.Transactions = append(blockInfo.Transactions, fmt.Sprintf("%v", t.Hash().String()))
	}

	return blockInfo
}

func (r *eth) GetTransactionByHash(txHash string) models.GetTransactionByHashResponse {
	ctx := context.Background()
	tx, _, err := r.client.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(r.chainId), big.NewInt(tx.GasFeeCap().Int64()))
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}

	var transaction models.GetTransactionByHashResponse
	transaction.TxHash = tx.Hash().String()
	transaction.From = msg.From().String()
	transaction.To = tx.To().String()
	transaction.Nonce = int64(tx.Nonce())
	transaction.Value = msg.Value().String()
	transaction.Data = fmt.Sprintf("0x%s", hex.EncodeToString(msg.Data()))
	receipt, err := r.client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		lg.Logging.ErrorLogging(err)
	}
	txLogs := make([]models.TransactionLog, 0)
	for _, itm := range receipt.Logs {
		var txLog models.TransactionLog
		txLog.Index = itm.Index
		txLog.Data = fmt.Sprintf("0x%s", hex.EncodeToString(itm.Data))
		txLogs = append(txLogs, txLog)
	}
	transaction.Logs = txLogs

	return transaction
}
