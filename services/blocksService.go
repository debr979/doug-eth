package services

import (
	"doug/models"
	msg "doug/utils/errMsg"
	"doug/utils/ether"
)

type blocks struct {
}

var Blocks blocks

func (r *blocks) GetBlockByCount(req models.GetBlockByCountRequest) interface{} {
	if req.Limit <= 0 {
		return msg.Messenger.Get(msg.ErrorParams)
	}

	var response models.GetBlockByCountResponse
	response.Blocks = ether.Eth.GetBlockInfoByCount(req.Limit)
	return response
}

func (r *blocks) GetBlockById(id int64) interface{} {
	return ether.Eth.GetBlockInfoById(id)
}
