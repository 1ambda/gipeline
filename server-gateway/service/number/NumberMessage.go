package number

import "github.com/a-trium/gipeline/server-gateway/service/common"

type InsertRequest struct {
	User string `json:"user"`
	// TODO: type alias
	Number int `json:"number"`
}

type InsertResponse struct {
	Message string `json:"message"`
	common.ErrResponse
}

type TotalResponse struct {
	Total int `json:"total"`
	common.ErrResponse
}
