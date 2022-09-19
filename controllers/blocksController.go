package controllers

import (
	"doug/models"
	"doug/services"
	msg "doug/utils/errMsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

type blocks struct {
	baseController
}

var Blocks blocks

func (r *blocks) GetBlockByCount(c *gin.Context) {

	var req models.GetBlockByCountRequest
	if err := r.GetBodyData(c, &req); err != nil {
		r.ResponseJSON(msg.Messenger.Get(msg.ErrorParams))
		return
	}

	r.ResponseJSON(services.Blocks.GetBlockByCount(req))
}

func (r *blocks) GetBlockById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.ResponseJSON(msg.Messenger.Get(msg.ErrorParams))
		return
	}

	r.ResponseJSONWithCtx(c, services.Blocks.GetBlockById(int64(id)))
}
