package controllers

import (
	msg "doug/utils/errMsg"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type baseController struct {
	Ctx *gin.Context
}

var Base baseController

func (b *baseController) GetBodyData(c *gin.Context, req interface{}) error {
	b.Ctx = c
	if err := b.Ctx.ShouldBind(req); err != nil {
		return err
	}

	return nil
}

func (b *baseController) GetHeaderData() error {
	return nil
}

func (b *baseController) GetFormFile(name string) (multipart.File, *multipart.FileHeader, error) {
	header, err := b.Ctx.FormFile(name)
	if err != nil {
		return nil, nil, err
	}
	file, err := header.Open()
	if err != nil {
		return nil, nil, err
	}

	return file, header, nil
}

func (b *baseController) ResponseJSON(result interface{}) {
	// Init Json Response Struct
	jsonResp := struct {
		ErrorCode int         `json:"errorCode"`
		ErrorMsg  string      `json:"errorMsg"`
		Data      interface{} `json:"data"`
	}{}

	var emptyData struct{}

	// Handle Data in Different Type
	switch value := result.(type) {
	case *msg.Msg:
		jsonResp.ErrorCode = value.Status
		jsonResp.ErrorMsg = value.Message
		jsonResp.Data = emptyData
	case string:
		var jsonData interface{}

		err := json.Unmarshal([]byte(value), &jsonData)

		if err != nil {
			jsonResp.ErrorCode = 5000
			jsonResp.ErrorMsg = fmt.Sprintf("Json Encoding Error: %s", err.Error())
			break
		}

		jsonResp.ErrorCode = 2000
		jsonResp.Data = jsonData
	default:
		jsonResp.ErrorCode = 2000
		jsonResp.Data = value
	}

	b.Ctx.JSONP(200, jsonResp)
	return
}

func (b *baseController) ResponseJSONWithCtx(c *gin.Context, result interface{}) {
	// Init Json Response Struct
	jsonResp := struct {
		ErrorCode int         `json:"errorCode"`
		ErrorMsg  string      `json:"errorMsg"`
		Data      interface{} `json:"data"`
	}{}

	// Handle Data in Different Type
	switch value := result.(type) {
	case *msg.Msg:
		jsonResp.ErrorCode = value.Status
		jsonResp.ErrorMsg = value.Message
	case string:
		var jsonData interface{}

		err := json.Unmarshal([]byte(value), &jsonData)

		if err != nil {
			jsonResp.ErrorCode = 5000
			jsonResp.ErrorMsg = fmt.Sprintf("Json Encoding Error: %s", err.Error())
			break
		}

		jsonResp.ErrorCode = 2000
		jsonResp.Data = jsonData
	default:
		jsonResp.ErrorCode = 2000
		jsonResp.Data = value
	}

	c.JSONP(200, jsonResp)
	return
}
