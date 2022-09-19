package models

// GetBlockByCountRequest GetBlockByCount
type GetBlockByCountRequest struct {
	Limit int `json:"limit" form:"limit" binding:"required"`
}

type BlockInfo struct {
	BlockNum   int64  `json:"block_num"`
	BlockHash  string `json:"block_hash"`
	BlockTime  int64  `json:"block_time"`
	ParentHash string `json:"parent_hash"`
}

type GetBlockByCountResponse struct {
	Blocks []BlockInfo `json:"blocks"`
}

// GetBlockByIdRequest GetBlockById
type GetBlockByIdRequest struct {
	Id int `json:"id" form:"id" binding:"required"`
}

type GetBlockByIdResponse struct {
	BlockNum     int64    `json:"block_num"`
	BlockHash    string   `json:"block_hash"`
	BlockTime    int64    `json:"block_time"`
	ParentHash   string   `json:"parent_hash"`
	Transactions []string `json:"transactions"`
}
